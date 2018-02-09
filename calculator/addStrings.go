package calculator

import (
	"strconv"
	"strings"
)

func add(numString string) int {
	delimiter := "\n"
	numString = strings.Replace(numString, delimiter, ",", -1)

	numbers := strings.Split(numString, ",")
	output := 0

	for _, number := range numbers {
		thisNumber, _ := strconv.Atoi(number)
		output = output + thisNumber
	}

	return output
}