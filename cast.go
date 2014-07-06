package t

import (
	"fmt"
	"strconv"
)

func to_str(v interface{}) string {
	if val, ok := v.(string); ok {
		return val
	}

	return fmt.Sprint(v)
}

func to_bool(v interface{}) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	return false
}

func to_int(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case int32:
		return int(val)
	case int16:
		return int(val)
	case int8:
		return int(val)
	case string:
		if i, err := strconv.ParseInt(val, 10, 32); err == nil {
			return int(i)
		}
	default:
		panic("T.Value type is not supported")
	}
	return 0
}

func to_int64(v interface{}) int64 {
	switch val := v.(type) {
	case int64:
		return val
	case string:
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return i
		}
	}
	return int64(to_int(v))
}

func to_uint(v interface{}) uint {
	switch val := v.(type) {
	case uint:
		return val
	case uint64:
		return uint(val)
	case uint32:
		return uint(val)
	case uint16:
		return uint(val)
	case uint8:
		return uint(val)
	case string:
		if i, err := strconv.ParseInt(val, 10, 32); err == nil {
			return uint(i)
		}
	default:
		panic("T.Value type is not supported")
	}
	return 0
}

func to_uint64(v interface{}) uint64 {
	switch val := v.(type) {
	case uint64:
		return val
	case string:
		if i, err := strconv.ParseUint(val, 10, 64); err == nil {
			return i
		}
	}
	return uint64(to_uint(v))
}
