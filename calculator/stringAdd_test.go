package calculator

import (
	"testing"
)

var addTests = []struct {
	input			string 	// testing input
	exOutput 	int			// expected output
}{
	{"1,2,3", 6},
	{"", 0},
	{"2112", 2112},
	{"1\n2,3", 6},
	{"2,3", 5},
	{"//;\n1;2", 3},
}

func TestAdd(t *testing.T) {
	for _,testCase := range addTests {
		actual := Add(testCase.input)
		if actual != testCase.exOutput {
			t.Errorf("Fail! For input: %s | Expected: %d | Received: %d", testCase.input, testCase.exOutput, actual)
		}
	}


}