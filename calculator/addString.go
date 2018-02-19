package calculator

import (
	"strconv"
	"strings"
)

func findDelim(delimString string) string {
	startDelim := 2
	endDelim := strings.Index(delimString, "\n")
	runes := []rune(delimString)
	safeSubString := string(runes[startDelim:endDelim])
	return safeSubString
}

func Add(numString string) int {
	output := 0
	delimiter := ","

	// test for new delim
	if strings.Index(numString, "\\") == 0 {
		baseDelim := findDelim(numString)
		endDelim := strings.Index(numString, "\n")

		numString = numString[endDelim:]
		numString = strings.Replace(numString, baseDelim, delimiter, -1)
	}

	// replace new line characters
	numString = strings.Replace(numString, "\n", delimiter, -1)

	// splt on delimiter
	numbers := strings.Split(numString, delimiter)

	// range over numbers summing the value of each element
	for _, number := range numbers {
		thisNumber, _ := strconv.Atoi(number)
		output = output + thisNumber
	}

	return output
}
