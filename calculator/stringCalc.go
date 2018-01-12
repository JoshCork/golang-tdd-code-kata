package calculator

import (
	"strconv"
	"strings"
)

func Add(numString string) int {
	stSlice := strings.Split(numString, ",")
	numCount := len(stSlice)

	addResult := 0

	if numCount == 1 {
		if stSlice[0] == "" {
			addResult = 0
		} else {
			addResult, _ = strconv.Atoi(stSlice[0])
		}
	} else {
		for _, num := range stSlice {
			convNum, _ := strconv.Atoi(num)
			addResult += convNum
	}
	}

	return addResult
}