package calculator

import (
	"testing"
)

// create table driven test struct

var addTests = []struct {
	input		string	// test input
	expOut	int			// expected output
}{
	{"", 0},
	{"2112", 2112},
	{"1,2,3", 6},
	{"1\n2,3", 6},
	{"\\;\n1;2;3", 6},
}

func TestAdd(t *testing.T){

	// iterate over table driven test struct
	for _,testCase := range addTests {
		actual := Add(testCase.input)
		// assert that output for given input meets expectations
		if actual != testCase.expOut {
			// raise error if that is not the case
			t.Errorf("Oh snap! Input: %s | Expected: %d | Actual: %d", testCase.input, testCase.expOut, actual)
		}
	}

}