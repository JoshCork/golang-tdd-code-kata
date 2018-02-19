package calculator

import (
	"testing"
)

// create table driven test struct

var addTestsPositive = []struct {
	input  string // test input
	expOut int    // expected output
}{
	{"", 0},
	{"2112", 2112},
	{"1,2,3", 6},
	{"1\n2,3", 6},
	{"\\;\n1;2;3", 6},
}

var addTestsNegative = []struct {
	input  string // test input
	expOut error  // expected output
}{
	{"-1", error{}},
}

// var addTestsNegative = []struct {
// 	input		string							// test input
// 	expOut	NegativeNumberErr{} // expected output
// }{
// 	{"-1", NegativeNumberErr{}},
// }

func TestAdd(t *testing.T) {

	// iterate over table driven test struct
	for _, testCase := range addTestsPositive {
		actual := Add(testCase.input)
		// assert that output for given input meets expectations
		if actual != testCase.expOut {
			// raise error if that is not the case
			t.Errorf("Oh snap! Input: %s | Expected: %d | Actual: %d", testCase.input, testCase.expOut, actual)
		}
	}

}

func TestAdd_Negative(t *testing.T) {

	// iterate over table driven test struct
	for _, testCase := range addTestsNegative {
		actual := Add(testCase.input)
		// assert that output for given input meets expectations
		if actual != testCase.expOut {
			// raise error if that is not the case
			t.Errorf("Oh snap! Input: %s | Expected: %d | Actual: %d", testCase.input, testCase.expOut, actual)
		}
	}

}
