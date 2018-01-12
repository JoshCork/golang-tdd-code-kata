package calculator

import (
	"strconv"
	"strings"
)


func Add(numString string) int {
	nums := strings.Split(numString, ",")
	rValue := 0

		for _, num := range nums {
			thisInt,_ := strconv.Atoi(num)
			rValue = rValue +  thisInt
		}

	return rValue
}