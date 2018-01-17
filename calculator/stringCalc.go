package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(nums string) int {

	fmt.Println(strings.Index(nums, "//"))
	fmt.Println(strings.HasPrefix(nums, "//"))

	if strings.HasPrefix(nums, "//") > -1 {
		// prefix := [element 1 of slace after doing a nums split string on \n after doing a trimprfix on \\]
		// fNums := replace element 2 of split string on nums \n replacing prefix with a comma

	} // else { fNums from line 20 }

	fNums := strings.Replace(nums, "\n", ",", -1)
	sum := 0
	numsSlice := strings.Split(fNums, ",")

	for _,num := range numsSlice {
		intNum,_ := strconv.Atoi(num)
		sum = sum + intNum
	}

	return sum

}