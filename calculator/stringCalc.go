package calculator

import (
	"strconv"
	"strings"
)

func Add(nums string) int {

	// need to parse input looking for "// followed by /"
	/*
		if the above is found:
			- I need to use the value inbetween those two indexes as the delimiter in my splitstr
			- I need trim that delimiter portion of the string off the front before I pass it into my splitter

		If the above is not found then I don't change anything and use the range below.
	*/



	r := strings.NewReplacer("\n", ",")
	formattedStrings := r.Replace(nums)
	numbers := strings.Split(formattedStrings, ",")

	returnValue := 0

	for _, sNum := range numbers {
		iNum, _ := strconv.Atoi(sNum)
		returnValue = returnValue + iNum
	}

	return returnValue

}