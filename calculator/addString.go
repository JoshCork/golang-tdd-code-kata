package calculator

import (
	"regexp"
	"strconv"
	"strings"
	"errors"
	"fmt"
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
			// leaving off here today.  took this down a different experimental route with a clean
			// string function. Need top pick up here tomorrow.

		}
	return safeSubString, endDelim
}

func cleanString(dirtyString string) string {
	delimString := dirtyString[1:strings.Index(dirtyString, "\n")]
	numString := dirtyString[strings.Index(dirtyString, "\n")+1:]

}

func test() {

	dirtyString := ""
	dirtyString = "\\***\n1***2***3***4"
	delimString := dirtyString[1:strings.Index(dirtyString, "\n")]
	numString := dirtyString[strings.Index(dirtyString, "\n")+1:]

	fmt.Printf("%q\n",delimString)
	fmt.Printf("%q\n",numString)
	fmt.Printf("%q\n",dirtyString)

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