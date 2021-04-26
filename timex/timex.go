package timex

import (
	"fmt"
	"regexp"
	"time"
)

const timeFmt = "2006-01-02 15:04:05"

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
	return time.Now().Format(timeFmt)
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
	loc, _ := time.LoadLocation("Asia/Shanghai")
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
	loc, _ := time.LoadLocation("Asia/Shanghai")
	y, m, d := time.Now().In(loc).AddDate(0, 0, -1).Date()
	start := fmt.Sprintf("%d-%02d-%02d 00:00:00", y, m, d)
	end := fmt.Sprintf("%d-%02d-%02d 23:59:59", y, m, d)
	return start, end
}

//判断当前时间时间是否在中间
func CompareTime(startDate string, endDate string) (bool, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var nowTime time.time = time.Now().In(loc)

	starTime, err := time.ParseInLocation(timeFmt, startDate, loc)
	if err != nil {
		return false, err
	}

	endTime, err := time.ParseInLocation(timeFmt, endDate, loc)

	if err != nil {
		return false, err
	}

	if nowtime.After(starTime) && nowtime.Before(endTime) {
		return true, nil
	}

	return false, nil

}

//判断时间格式是否正确 - xxxx-xx-xx 00:00:00
func CheckTimeDate(date string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$`)
	return reg.Match([]byte(date))
}
