package util

import (
	"fmt"
	"time"
)

// SubStr safe substring function for both ascii and unicode strings
func SubStr(s string, start, end int) string {
	runes := []rune(s)
	size := len(runes)
	if start > size {
		return ""
	}
	if end > size {
		end = size
	}
	return string(runes[start:end])
}

// LogStr return a string for logging, truncate with size 500
func LogStr(ins interface{}) string {
	size := 1000
	switch val := ins.(type) {
	case string:
		return SubStr(val, 0, size)
	case []byte:
		if len(val) > size {
			val = val[:size]
		}
		return string(val)
	case int64, int32, int, bool, int8, uint, uint32, uint64:
		return fmt.Sprintf("%d", val)
	default:
		return SubStr(JSONStr(val), 0, size)
	}

}

// StatusMsg convert string to *string in order to set the status msg.
func StatusMsg(v string) *string {
	return &v
}

// TimeStringToGoTime 时间格式字符串转换
func TimeStringToGoTime(tm string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", tm, time.Local)
	if nil == err && !t.IsZero() {
		return t
	}
	return t
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
