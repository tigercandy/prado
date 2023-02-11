package utils

import (
	"strconv"
	"strings"
)

func String2Int(s string) (int, error) {
	return strconv.Atoi(s)
}

func IdsStr2IdsInt(ids_key string) []int {
	res := make([]int, 0)
	ids := strings.Split(ids_key, ",")
	for i := 0; i < len(ids); i++ {
		id, _ := String2Int(ids[i])
		res = append(res, id)
	}
	return res
}
