package calculator

import (
	"strconv"
	"strings"
)

func helpSubstring(myString string) string {
	delimiterEnd := strings.Index(myString, "\n")

}

func Add(numString string) int {
	output := 0
	delimiter := ","

	if strings.Index(numString, "//") = 0 {
		delimiter =
	}

	numString = strings.Replace(numString, delimiter, ",", -1)

	numbers := strings.Split(numString, ",")

	for _,number := range	numbers {
		thisNubmer,_ := strconv.Atoi(number)
		output = output + thisNubmer
	}

	return output

}