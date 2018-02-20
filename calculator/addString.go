package calculator

import (
	"strconv"
	"strings"
	"errors"
)

func findDelim (delimString string) (string, int) {
	delimStart := 1
	endDelim := strings.Index(delimString, "\n")


	runes := []rune(delimString)
	safeSubString := string(runes[delimStart:endDelim])
		if strings.Index(safeSubString, "[") == 0 {
			safeSubString = strings.Replace(safeSubString, "[", "", -1)
			safeSubString = strings.Replace(safeSubString, "]", "", -1)
		}
	return safeSubString, endDelim
}

func Add(numString string) (int, error) {
	output := 0
	delimiter := ","
	theNegatives := ""

	if strings.Index(numString, "\\") == 0 {
		baseDelim, endDelim := findDelim(numString)
		numString = numString[endDelim:]
		numString = strings.Replace(numString, baseDelim, delimiter, -1)
	}

	numString = strings.Replace(numString, "\n", delimiter, -1)
	numbers := strings.Split(numString, delimiter)

	for _,number := range numbers {
		theNumber,_ := strconv.Atoi(number)
		if theNumber < 0 {
			theNegatives = theNegatives + "," + number
		} else if theNumber > 1000 {
			// do nothing, ignore this number
		} else {
			output = output + theNumber
		}
	}

	if theNegatives != "" {
		return -1, errors.New("Cannot handle the negativity " + theNegatives)
	}
	return output, nil
}