package calculator

import (
	"strconv"
	"strings"
)

// Finds substring for delimiter
func findDelim(delimNumString string) string {
	endDelim := strings.Index(delimNumString, "\n")
	runes := []rune(delimNumString)
	safeSubstring := string(runes[2:endDelim])

	return safeSubstring
}

// Need a truncate left helper function
func truncateLeft(delimNumString string) string {
	endDelim := strings.Index(delimNumString, "\n")
	runes := []rune(delimNumString)
	safeSubstring := string(runes[endDelim:])

	return safeSubstring
}

// Add summarize a set of delimited numbers in string format
// Calling Add with a negative number will throw an exception “negatives not allowed” and the negative that was passed
// If there are multiple negatives, show all of them in the exception message
func Add(numString string) int {

	output := 0
	delimiter := ","

	if strings.Index(numString, "\\") == 0 {
		baseDelim := findDelim(numString)
		numString = truncateLeft(numString)
		numString = strings.Replace(numString, baseDelim, delimiter, -1)
	}

	// convert all "\n" to commas
	numString = strings.Replace(numString, "\n", ",", -1)

	// convert string to slice and iterate over slice
	intStrings := strings.Split(numString, delimiter)
	for _,num := range intStrings {
		// convert element of type string to integer
		thisNum,_ := strconv.Atoi(num)

		// keep running total of elements
		output = output + thisNum
	}


	// return running total of elements
	return output
}