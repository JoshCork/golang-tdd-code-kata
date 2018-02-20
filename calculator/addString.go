package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func findDelim(delimString string) string {
	delimStart := 2
	delimEnd := strings.Index(delimString, "\n")
	runes := []rune(delimString)
	safeSubString := string(runes[delimStart:delimEnd])

	return safeSubString
}

func Add(numString string) (int, error) {
	output := 0
	delimiter := ","
	negatives := ""

	if strings.Index(numString, "\\") == 0 {
		baseDelim := findDelim(numString)
		delimEnd := strings.Index(numString, "\n")
		numString = numString[delimEnd:]
		numString = strings.Replace(numString, baseDelim, delimiter, -1)
	}

	numString = strings.Replace(numString, "\n", delimiter, -1)

	numbers := strings.Split(numString, delimiter)
	for _, number := range numbers {
		thisNumber, _ := strconv.Atoi(number)
		if thisNumber < 0 {
			negatives = negatives + "," + number
		} else if thisNumber > 1000 {
			// do nothing
		} else {
			output = output + thisNumber
		}
	}

	if negatives != "" {
		errText := "Cannot handle the negativity" + negatives
		return -1, errors.New(errText)
	}
	return output, nil

}
