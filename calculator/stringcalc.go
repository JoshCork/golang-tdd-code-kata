package calculator

import (
	"fmt"
	"strings"
	"strconv"
)

func Add(stringNum string) int {
	numbers := strings.Split(stringNum, ",")

	fmt.Println(len(numbers))
	fmt.Printf("numbers[0]:%v", numbers[0])

	if len(numbers) == 0 {
		switch numbers[0] {
		case "":
			return 0
		default:
			i, _ := strconv.Atoi(numbers[0])

			return i
		}
	}

	return 1

}