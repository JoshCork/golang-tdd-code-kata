package calculator

import (
	"testing"
)

var addTests = []struct {
	inString				string 	//input for test
	expectedOutput	int			// expected out putput for given test input
}{
	{"0",0},
	{"",0},
	{"1,2,3",6},
	{"2112", 2112},
	{"1\n2,3", 6},
	{"//;\n1;2", 3},
}

func TestAdd(t *testing.T) {

	for _,testCase := range addTests {
		actual := add(testCase.inString)
		if actual != testCase.expectedOutput {
			t.Errorf("\nOh snap! Input: %s | Expected Output: %d | Actual Output: %d", testCase.inString, testCase.expectedOutput, actual)
		}
	}

}