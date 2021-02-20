package gconv

import "fmt"

func String(v interface{}) string {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case string:
		return fmt.Sprintf("%s", v)
	default:
		return ""
	}
}
