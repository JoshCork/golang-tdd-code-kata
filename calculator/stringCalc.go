package calculator

import (
	"strconv"
	"fmt"
	"strings"
)

func Add(numString string) int {

	numSlice := strings.Split(numString, ",")
	returnVal := 0




	fmt.Printf("\ninput value was: %s\n", numString)
	fmt.Printf("element 0 of numSlice: %s\n", numSlice[0])
	fmt.Printf("lenth of input is: %d\n", len(numSlice))

	if len(numSlice) == 1 {
		if numSlice[0] == "" {
			fmt.Print("element zero was an empty string\n")
		} else {
		for i := 0; i <= len(numSlice); i++ {
			currentIntVal, _ := strconv.Atoi(numSlice[i])
			returnVal = returnVal + currentIntVal
		}
		}
	}	else {
		for i := 0; i < len(numSlice); i++ {
			currentIntVal, _ := strconv.Atoi(numSlice[i])
			returnVal = returnVal + currentIntVal
		}
	}

	return returnVal

}