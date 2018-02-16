package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

// helpSubstring is a function used to clip substrings from a long string
func helpSubstring(fullString string, startSub int, endSub int)  string {
	// convert to rune slice
	runes := []rune(fullString)
	safeSubstring := string(runes[startSub:endSub])

	return safeSubstring
}

// helpTrimLeft is a helper function that truncates the left portion of a string
func helpTrimLeft(fullString string, startSub int)  string {
	// convert to rune slice
	runes := []rune(fullString)
	safeSubstring := string(runes[startSub:])

	return safeSubstring
}

// Add takes a string of delimited numbers and perform a sum on those numbers
// it defaults to a comma delmited set of numbers but also allows for other delimiters
// to be baked into the string itself as a prefix
func Add(numString string) int {
	output := 0
	delimiter := ","

	if strings.Index(numString, "//") == 0 {
		// find end delimiter
		endDelim := strings.Index(numString, "\n")

		delimiter = helpSubstring(numString, 2, endDelim)
		fmt.Printf("\nDemlimiter: %s", delimiter)

		numString = helpTrimLeft(numString, endDelim)
		fmt.Printf("\nnumString: %s", numString)

	}

	// swap the delmiter out for commas
	numString = strings.Replace(numString, delimiter, ",", -1)
	// swap out any newline breaks with commas
	numString = strings.Replace(numString, "\n", ",", -1)

	//split on comma
	numbers := strings.Split(numString, ",")

	// convert each element to a number and sum
	for _,number := range	numbers {
		thisNubmer,_ := strconv.Atoi(number)
		output = output + thisNubmer
	}

	// return the summarized numbers
	return output

}