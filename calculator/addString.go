package calculator

import (
	"strconv"
	"strings"
)

// helper func to find new string
func findDelim(delimString string) string {
	// hard coded, this is part of the problem statment/requirements
	delimStart := 2
	runes := []rune(delimString)
	endDelim := strings.Index(delimString, "\n")

	safeSubString := string(runes[delimStart:endDelim])

	return safeSubString
}

func Add(numString string) int {

	output := 0
	delimiter := ","

	// test for delimiter
	if strings.Index(numString, "\\") == 0 {
		// find delimiter
		baseDelim := findDelim(numString)
		endDelim := strings.Index(numString, "\n")

		// truncate leading meatadata
		numString = numString[endDelim:]

		// replace delimiter
		numString = strings.Replace(numString, baseDelim, delimiter, -1)

	}

	// replace all new lines with commas
	numString = strings.Replace(numString, "\n", delimiter, -1)

	// split on delimiter
	numbers := strings.Split(numString, delimiter)

	for _, number := range numbers {
		thisNumber, _ := strconv.Atoi(number)
		output = output + thisNumber
	}

	// range over split and sum the total of each element

	return output
}
