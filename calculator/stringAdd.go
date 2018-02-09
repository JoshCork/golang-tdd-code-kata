package calculator

import (
	"strconv"
	"strings"
)


func Add(numString string) int {
	splitNums := strings.Split(numString, ",")
	sumValue := 0
	delimiter := ","

	if strings.Index(numString,"//") == 0 {
		endDelimiter := strings.Index(numString, "\n")
		delimiter =
	}

	for _, num := range splitNums {
		thisNum,_ := strconv.Atoi(num)
		sumValue = sumValue + thisNum
	}

	return sumValue
}