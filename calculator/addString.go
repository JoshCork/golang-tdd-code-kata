package calculator

import (
	"regexp"
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

			marker, _ := regexp.Compile(`\[(.*)\]`)
			bracketDelims := marker.FindAllString(safeSubString, -1)
			delims := []string

			for _,delim := range bracketDelims {
				delim = strings.Replace(safeSubString, "[", "", -1)
				delim = strings.Replace(safeSubString, "]", "", -1)
				delims = append(delims, delim)
			}
			// leaving off here today.  I need to map each delim in delims to a comma
			// and then pass back that map and execute the map on the actual string to be added
			// doing a replace on each using the map.
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