package timex

import (
	"fmt"
	"time"
)

const timeFmt = "2006-01-02 15:04:05"
const timeZone = "Asia/Shanghai"

//获取当前的时间戳
func NowTimeStamp() int64 {
	loc, _ := time.LoadLocation(timeZone)
	return time.Now().In(loc).Unix()
}

//解析当前日期 转换成 时间
func ParseStrTime(timeStr string) (time.Time, error) {
	t, err := time.ParseInLocation(timeFmt, timeStr, time.Local)
	return t, err
}

func ToRFC3339(now time.Time) string {
	return now.Format(timeFmt)
}

//获取当前时间
func StrNow() string {
	return time.Now().Local().Format(timeFmt)
}

//格式化时间格式
func ConvertRFC3339TimeFormat(t string) (string, error) {
	tmp, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return t, nil
	}
	return tmp.Format(timeFmt), nil
}

//获取相差时间（秒）
func GetSecondDiffer(startTime, endTime string) int64 {
	var second int64
	loc, _ := time.LoadLocation(timeZone)
	t1, err := time.ParseInLocation(timeFmt, startTime, loc)
	t2, err := time.ParseInLocation(timeFmt, endTime, loc)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		second = diff
		return second
	} else {
		return second
	}
}

//格式化时间戳
func GetTimeUnixToTimeStr(unix int) string {
	tm := time.Unix(int64(unix), 0)
	return tm.Format(timeFmt)
}

//Yesterday 返回昨天 起始、截止时间 xxxx-xx-xx 00:00:00 xx-xx-xx 23:59:59
func Yesterday() (string, string) {
	loc, _ := time.LoadLocation(timeZone)
	y, m, d := time.Now().In(loc).AddDate(0, 0, -1).Date()
	start := fmt.Sprintf("%d-%02d-%02d 00:00:00", y, m, d)
	end := fmt.Sprintf("%d-%02d-%02d 23:59:59", y, m, d)
	return start, end
}

//判断当前时间时间是否在中间
func CompareTime(startDate string, endDate string) (bool, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var nowTime time.Time = time.Now().In(loc)

	starTime, err := time.ParseInLocation(timeFmt, startDate, loc)
	if err != nil {
		return false, err
	}

	endTime, err := time.ParseInLocation(timeFmt, endDate, loc)

	if err != nil {
		return false, err
	}

	if nowTime.After(starTime) && nowTime.Before(endTime) {
		return true, nil
	}

	return false, nil

}

//compare 比较时间
//startDate 开始时间  endDate 结束时间 与当前时间比较
//返回 1 小于开始时间  2 大于开时间 && 小于 结束时间  3 大于结束时间 0 时间解析错误
func CompareDate(startDate string, endDate string, layout string) (res int, startTime time.Time, endTime time.Time) {
	res = 0
	if startDate == "" || endDate == "" {
		res = 0
		return
	}

	loc, _ := time.LoadLocation(timeZone)
	var nowtime time.Time = time.Now().In(loc)
	//var startTimt
	var err error

	if startTime, err = time.ParseInLocation(layout, startDate, loc); err != nil {
		res = 0
		return
	}

	if endTime, err = time.ParseInLocation(layout, endDate, loc); err != nil {
		res = 0
		return
	}

	if startTime.After(endTime) {
		res = 0
		return
	}

	if nowtime.Before(startTime) {
		res = 1
		return
	} else if nowtime.After(startTime) && nowtime.Before(endTime) {
		res = 2
		return
	} else if nowtime.After(endTime) {
		res = 3
		return
	}
	return
}