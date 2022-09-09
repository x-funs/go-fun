package fun

import (
	"time"

	"github.com/x-funs/go-fun/strtotime"
)

// Timestamp 返回当前时间的 Unix 时间戳。
// 默认返回秒级, 支持 Timestamp(true) 返回毫秒级
func Timestamp(millis ...any) int64 {
	l := len(millis)
	switch l {
	case 0:
		return unixTimestamp()
	case 1:
		switch v := millis[0].(type) {
		case bool:
			if v {
				return unixMilliTimestamp()
			}
		}
	}

	return unixTimestamp()
}

// unixTimestamp 返回当前时间的 Unix 秒级时间戳
func unixTimestamp() int64 {
	return time.Now().Unix()
}

// unixMilliTimestamp 返回当前时间的 Unix 毫秒级时间戳
func unixMilliTimestamp() int64 {
	return time.Now().UnixMilli()
}

// Date 返回格式化后的日期时间字符串。
// 支持 Date()、Date(unixstamp)、Date(layout)、Date(layout, unixstamp)
func Date(layouts ...any) string {
	l := len(layouts)

	switch l {
	case 0:
		return dateByDefault()
	case 1:
		switch v := layouts[0].(type) {
		case string:
			return dateByPattern(ToString(v))
		case int, int8, int16, int32, int64:
			return dateByPatternAndTime("", ToInt64(v))
		}
	case 2:
		switch layouts[0].(type) {
		case string:
			switch v := layouts[1].(type) {
			case int, int8, int16, int32, int64:
				return dateByPatternAndTime(ToString(layouts[0]), ToInt64(v))
			}
		}
	}

	return ""
}

// dateByDefault 返回默认 layout 格式化后的日期时间字符串
func dateByDefault() string {
	now := time.Now()
	return now.Format(DatetimePattern)
}

// dateByPattern 返回指定 layout 格式化后的日期时间字符串
func dateByPattern(layout string) string {
	now := time.Now()

	if Blank(layout) {
		return now.Format(DatetimePattern)
	} else {
		return now.Format(layout)
	}
}

// dateByPatternAndTime 返回指定时间戳、指定 layout 格式化后的日期时间字符串
func dateByPatternAndTime(layout string, timeStamp int64) string {
	if timeStamp < 0 {
		timeStamp = 0
	}
	uTime := time.Unix(timeStamp, 0)

	if Blank(layout) {
		return uTime.Format(DatetimePattern)
	} else {
		return uTime.Format(layout)
	}
}

// StrToTime 日期时间字符串转时间戳
// 支持 StrToTime()、StrToTime(string)、StrToTime(string, int64)
func StrToTime(args ...any) int64 {
	l := len(args)

	switch l {
	case 0:
		return Timestamp()
	case 1:
		exp := ToString(args[0])
		if !Blank(exp) {
			v, err := strtotime.Parse(exp, Timestamp())
			if err == nil {
				return v
			}
		}
	case 2:
		exp := ToString(args[0])
		if !Blank(exp) {
			timeStamp := ToInt64(args[1])
			if timeStamp > 0 {
				v, err := strtotime.Parse(exp, timeStamp)
				if err == nil {
					return v
				}
			}
		}
	}

	return 0
}
