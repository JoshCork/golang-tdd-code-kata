package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func findDelim(delimString string) string {
	startDelim := 2
	endDelim := strings.Index(delimString,"\n")

	runes := []rune(delimString)
	safeSubString := string(runes[startDelim:endDelim])

	return safeSubString
}

func Add(numString string) (int, error) {
	output := 0
	delimiter := ","
	errMessage := ""

	if strings.Index(numString, "\\") == 0 {
		baseDelim := findDelim(numString)
		endDelim := strings.Index(numString,"\n")
		numString = numString[endDelim:]
		numString = strings.Replace(numString, baseDelim, delimiter, -1)
	}

	numString = strings.Replace(numString, "\n", delimiter, -1)

	numbers := strings.Split(numString, delimiter)

		for _,number := range numbers {
		theNumber,_ := strconv.Atoi(number)
		if theNumber < 0 {
			errMessage = errMessage + "," + number
		} else {
			output = output + theNumber
		}
	}

	if errMessage != "" {
		return -1, errors.New("Cannot handle the negativity: " + errMessage)
	} else {
		return output, nil
	}
}