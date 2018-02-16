package calculator

import (
	"testing"
)

var addTests = []struct {
	inputString	string
	expectedOut	int
}{
	{"", 0},
	{"2112", 2112},
	{"2,3", 5},
	{"1\n2,3", 6},
	{"//;\n1;2", 3},
	{"//||\n1||2||4||2", 9},
}

func TestAdd(t *testing.T) {
	for _,testCase := range addTests {
		actual := Add(testCase.inputString)

		if actual != testCase.expectedOut {
			t.Errorf("Oh snap!! Input: %s | Expected Ouput: %d | Actual Ouput: %d", testCase.inputString, testCase.expectedOut, actual)
		}
	}
}