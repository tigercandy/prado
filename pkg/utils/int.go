package utils

import "strconv"

func Int2String(i int) string {
	return strconv.Itoa(i)
}

func Int642String(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int642Int(i int64) int {
	return int(i)
}
