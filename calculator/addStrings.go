package calculator

import (
	"strconv"
	"strings"
)

func Add(numString string) int {
	output := 0
	delimiter := "\n"

	numString = strings.Replace(numString, delimiter, ",", -1)

	numbers := strings.Split(numString, ",")

	for _,number := range	numbers {
		thisNubmer,_ := strconv.Atoi(number)
		output = output + thisNubmer
	}

	return output

}