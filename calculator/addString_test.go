package calculator

import (
	"testing"
)

var addTests = []struct {
	input  string // input string for numbers to be summed
	output int    // expected output from the function
	theErr string // error string returned
}{
	{"", 0, ""},
	{"1,2,3", 6, ""},
	{"2112", 2112, ""},
	{"1\n2,3", 6, ""},
	{"\\;\n1;2;3", 6, ""},
	{"\\|\n1|2|3|4", 10, ""},
	{"-1", -1, "can't work with negative numbers ,-1"},
	{"-1,-2,-3", -1, "can't work with negative numbers ,-1,-2,-3"},
}

func TestAdd(t *testing.T) {
	for _, testCase := range addTests {
		actual, err := Add(testCase.input)
		if err != nil {
			if err.Error() != testCase.theErr {
				t.Errorf("Oh Snap! For Input: %s | Expecting error: %s | Actual Error: %s", testCase.input, testCase.theErr, err.Error())
			}
		} else {
			if actual != testCase.output {
				t.Errorf("Oh Snap! Input: %s | Output: %d | Actual: %d", testCase.input, testCase.output, actual)
			}
		}

	}

}
