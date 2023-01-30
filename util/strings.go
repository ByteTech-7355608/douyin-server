package util

import "fmt"

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
