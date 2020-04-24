package helpers

import (
	"strconv"
	"strings"
)

func FormatArrayParams(params string, sep string) (arr []int) {
	if params == "" {
		arr = []int{}
	} else {
		strArray := strings.Split(params, sep)
		for _, s := range strArray {
			v, err := strconv.Atoi(s)
			if err != nil {
				arr = append(arr, v)
			}
		}
	}
	return
}