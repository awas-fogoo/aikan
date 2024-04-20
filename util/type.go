package util

import "strconv"

// StringToInt 字符串转int
func StringToInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return res
}

// StringToUint 字符串转uint
func StringToUint(v string) uint {
	res, err := strconv.Atoi(v)
	if err != nil {
		return uint(0)
	}
	return uint(res)
}

// StringToInt8 字符串转int8
func StringToInt8(v string) int8 {
	res, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return int8(0)
	}
	return int8(res)
}

// StringToUint8 字符串转uint8
func StringToUint8(v string) uint8 {
	res, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(res)
}

// StringToInt64 字符串转int64
func StringToInt64(v string) int64 {
	res, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return int64(0)
	}
	return res
}

// StringToUint64WithDecimal 字符串转uint64
func StringToUint64WithDecimal(v string) uint64 {
	res, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0
	}
	return uint64(res)
}
