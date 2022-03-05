package delayer

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Job struct {
	Topic string `json:"topic" msgpack:"1"`
	Id    string `json:"id" msgpack:"2"`    // job唯一标识ID
	Delay int64  `json:"delay" msgpack:"3"` // 延迟时间, unix时间戳
	TTR   int64  `json:"ttr" msgpack:"4"`
	Body  Body   `json:"body" msgpack:"5"`
}

type Body struct {
	BodyId   int64  `json:"body_id" msgpack:"1"`
	BodyDesc string `json:"body_desc" msgpack:"2"`
}

type Delayer struct {
	Cache          *redis.Client
	Timers         []*time.Ticker
	bucketNameChan <-chan string
	popHandle      func(string, []Body) error
	Ttr            int64
	topic          string
	MaxGet         int64
}

// BucketItem bucket中的元素
type BucketItem struct {
	timestamp int64
	jobId     string
}

func NewDelayer(cache *redis.Client, topic string) *Delayer {
	return &Delayer{
		Cache:  cache,
		topic:  topic,
		Ttr:    3,
		MaxGet: 20,
	}
}

// 初始化定时器
func (dey *Delayer) InitProducer() {
	dey.bucketNameChan = dey.generateBucketName()
}

//初始化消费者
func (dey *Delayer) InitConsumer() {
	dey.Timers = make([]*time.Ticker, 5)
	var bucketName string
	for i := 0; i < 5; i++ {
		dey.Timers[i] = time.NewTicker(1 * time.Second)
		bucketName = fmt.Sprintf("vp_bucket_%d", i+1)
		go dey.waitTicker(dey.Timers[i], bucketName)
	}

	dey.PopTaskTicker()
}

func (dey *Delayer) PopHandle(popHandle func(string, []Body) error) {
	dey.popHandle = popHandle
}

func (dey *Delayer) waitTicker(timer *time.Ticker, bucketName string) {
	for {
		select {
		case t := <-timer.C:
			dey.tickHandler(t, bucketName)
		}
	}
}

// 扫描bucket, 取出延迟时间小于当前时间的Job
func (dey *Delayer) tickHandler(t time.Time, bucketName string) {

	for {
		bucketItem, err := dey.getFromBucket(bucketName)
		if err != nil {
			logx.Infof("扫描bucket错误#bucket-%s#%s", bucketName, err.Error())
			return
		}

		// 集合为空
		if bucketItem == nil {
			return
		}

		// 延迟时间未到
		if bucketItem.timestamp > t.Unix() {
			return
		}

		// 延迟时间小于等于当前时间, 取出Job元信息并放入ready queue
		job, err := dey.getJob(bucketItem.jobId)

		if err != nil && err != redis.Nil {
			logx.Infof("获取Job元信息失败#bucket-%s#%s", bucketName, err.Error())
			continue
		}

		// job元信息不存在, 从bucket中删除
		if job == nil || err == redis.Nil {
			dey.removeFromBucket(bucketName, bucketItem.jobId)
			continue
		}

		// 再次确认元信息中delay是否小于等于当前时间

		if job.Delay > t.Unix() {
			// 从bucket中删除旧的jobId
			dey.removeFromBucket(bucketName, bucketItem.jobId)
			// 重新计算delay时间并放入bucket中
			dey.pushToBucket(<-dey.bucketNameChan, job.Delay, bucketItem.jobId)
			continue
		}

		err = dey.pushToReadyQueue(job.Topic, bucketItem.jobId)
		//dey.ready <- bucketItem.jobId
		if err != nil {
			logx.Infof("JobId放入ready queue失败#bucket-%s#job-%+v#%s",
				bucketName, job, err.Error())
			continue
		}

		// 从bucket中删除
		dey.removeFromBucket(bucketName, bucketItem.jobId)
	}
}

//添加任务
func (dey *Delayer) PushTask(jobId string, delay int64, bodyId int64, bodyDesc string) error {
	logx.Infof("延时队列添加任务：jobId:%s,delay:%d,bodyId:%d,bodyDesc:%s", jobId, delay, bodyId, bodyDesc)
	if jobId == "" {
		jobId = GenUniqueID()
	}

	if delay <= 0 || delay > (1<<31) {
		return errors.New("delay is error")
	}

	var job Job = Job{
		Id:    jobId,
		Delay: time.Now().Unix() + delay,
		TTR:   dey.Ttr,
		Topic: dey.topic,
		Body:  Body{BodyId: bodyId, BodyDesc: bodyDesc},
	}

	return dey.pushJob(job)
}

func (dey *Delayer) pushJob(job Job) error {
	if job.Id == "" || job.Topic == "" || job.Delay < 0 || job.TTR <= 0 {
		return errors.New("invalid job")
	}

	err := dey.putJob(job.Id, job)
	if err != nil {
		logx.Infof("添加job到job pool失败#job-%+v#%s", job, err.Error())
		return err
	}

	err = dey.pushToBucket(<-dey.bucketNameChan, job.Delay, job.Id)
	if err != nil {
		logx.Infof("添加job到bucket失败#job-%+v#%s", job, err.Error())
		return err
	}

	return nil
}

// 添加Job
func (dey *Delayer) putJob(key string, job Job) error {
	value, err := msgpack.Marshal(job)
	if err != nil {
		return err
	}
	return dey.Cache.Set(key, string(value), time.Minute*60).Err()
}

func (dey *Delayer) getFromBucket(key string) (*BucketItem, error) {

	value, err := dey.Cache.ZRangeWithScores(key, 0, 0).Result()
	if err != nil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, nil
	}

	item := &BucketItem{}
	item.timestamp = int64(value[0].Score)
	item.jobId = value[0].Member.(string)

	return item, nil
}

// 添加JobId到bucket中
func (dey *Delayer) pushToBucket(key string, timestamp int64, jobId string) error {

	memeber := &redis.Z{
		Score:  float64(timestamp),
		Member: jobId,
	}

	return dey.Cache.ZAdd(key, memeber).Err()
}

func (dey *Delayer) pushToReadyQueue(queueName string, jobId string) error {
	queueName = fmt.Sprintf("vp_queue_%s", queueName)
	return dey.Cache.RPush(queueName, jobId).Err()
}

// 从bucket中删除JobId
func (dey *Delayer) removeFromBucket(bucket string, jobId string) error {
	logx.Infof("从bucket中删除JobId,bucket:%s,jobId:%s，开始删除", bucket, jobId)
	err := dey.Cache.ZRem(bucket, jobId).Err()
	if err != nil {
		logx.Infof("从bucket中删除JobId,bucket:%s,jobId:%s,错误信息为：%s", bucket, jobId, err.Error())
	}
	return err
}

// 获取Job
func (dey *Delayer) getJob(key string) (*Job, error) {
	value, err := dey.Cache.Get(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if value == "" || err == redis.Nil {
		return nil, nil
	}

	job := &Job{}
	err = msgpack.Unmarshal([]byte(value), job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// 轮询获取bucket名称, 使job分布到不同bucket中, 提高扫描速度
func (dey *Delayer) generateBucketName() <-chan string {
	c := make(chan string)
	go func() {
		i := 1
		for {
			c <- fmt.Sprintf("vp_bucket_%d", i)
			if i >= 5 {
				i = 1
			} else {
				i++
			}
		}
	}()

	return c
}

// Pop 轮询获取Job
func (dey *Delayer) Pop(topics []string) ([]Body, error) {
	jobArr, err := dey.blockPopFromReadyQueue(topics, 178)
	if err != nil {
		return nil, err
	}

	// 队列为空
	if len(jobArr) == 0 {
		return nil, nil
	}

	var jobObjArr []Body = make([]Body, 0)
	for _, jobId := range jobArr {
		job, errs := dey.getJob(jobId)
		if errs != nil { //查询错误
			continue
		}

		if job == nil { //没有查询到
			continue
		}

		jobObjArr = append(jobObjArr, job.Body)
	}

	if errdel := dey.RemoveJob(jobArr); errdel != nil {
		logx.Info("删除key失败：", errdel.Error())
	}

	return jobObjArr, nil

}

// 从队列中阻塞获取JobId
func (dey *Delayer) blockPopFromReadyQueue(queues []string, timeout int) ([]string, error) {
	//var args []string
	var queue string = fmt.Sprintf("vp_queue_%s", queues[0])
	value, err := dey.Cache.LRange(queue, 0, dey.MaxGet).Result()

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(value) == 0 {
		return nil, nil
	}

	//删除数据
	if err = dey.Cache.LTrim(queue, int64(len(value)), -1).Err(); err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return value, nil
}

// 删除Job
func (dey *Delayer) RemoveJob(key []string) error {
	return dey.Cache.Del(key...).Err()
}

//轮训执行数据
func (dey *Delayer) PopTaskTicker() {

	t := time.NewTimer(2 * time.Second)
	go func() {
		for {
			select {
			case <-t.C:
				dey.popTask()
				t.Reset(2 * time.Second)
			}

		}

		defer t.Stop()

	}()
}

//轮询队列获取任务
func (dey *Delayer) popTask() {
	//logx.Info("开始循环取出队列中的数据")
	var topics []string = []string{dey.topic}
	jobInfo, err := dey.Pop(topics)
	if err != nil || jobInfo == nil || len(jobInfo) == 0 {
		return
	}
	dey.popHandle(dey.topic, jobInfo) //执行会掉函数
}
