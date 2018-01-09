package calculator

import (
	"strconv"
	"strings"
)

func Add(numString string) int {

	numsToAdd := strings.Split(numString, ",")

	returnValue := 0

	if len(numsToAdd) > 0 {
		if numsToAdd[0] == "" {
			return 0
		}
		returnValue, _ := strconv.Atoi(numsToAdd[0])

		return returnValue
	}

	return 100

}