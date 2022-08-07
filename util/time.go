package util

import "time"

// GetCurrentTime 获取当前时间  返回time.Time格式
func GetCurrentTime() time.Time {
	return time.Now() //打印结果 2021-04-11+12:52:52.794351777++0800+C
}

// GetCurrentTimestamp 获取当前时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix() //单位秒,打印结果//1531292871
	//return time.Now().UnixNano() //单位纳秒,打印结果//1531888244752784461
}

// GetCurrentDatetime 获取当前时间的字符串格式
func GetCurrentDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05") //打印结果 2021-11-01 13:24:25

}

// TimestampToDatetime 时间戳转时间字符串  int64->string
func TimestampToDatetime(tm int64) string {
	//timeUnix := time.Now().Unix()   //已知的时间戳
	timeUnix := tm
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05") //打印结果：2017-04-11 13:30:39

}

// TimestampToTime 时间戳转时间  int64->time
func TimestampToTime(tm int64) time.Time {
	timeUnix := TimestampToDatetime(tm)
	return DatetimeToTime(timeUnix) //2021-04-11+12:52:52.794351777++0800+C

}

// DatetimeToTime 字符串时间转时间戳  string->Time
func DatetimeToTime(formatTimeStr string) (tm time.Time) {
	//formatTimeStr:="2017-04-11 13:33:37"
	formatTime, _ := time.Parse("2006-01-02 15:04:05", formatTimeStr)

	return formatTime
	//打印结果：2017-04-11 13:33:37 +0000 UTC

}

// DatetimeToTimestamp 字符串时间转时间戳  string->int64
func DatetimeToTimestamp(formatTimeStr string) (tm int64) {
	//formatTimeStr:="2017-04-11 13:33:37"
	formatTime, _ := time.Parse("2006-01-02 15:04:05", formatTimeStr)

	return formatTime.Unix()

	//打印结果：1531888244752784461

}
