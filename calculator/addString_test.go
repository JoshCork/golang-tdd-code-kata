package calculator

import "testing"

var addTests = []struct {
	input  string // input string for numbers to be summed
	output int    // expected output from the function
}{
	{"", 0},
	{"1,2,3", 6},
	{"2112", 2112},
	{"1\n2,3", 6},
	{"\\;\n1;2;3", 6},
	{"\\|\n1|2|3|4", 10},
}

func TestAdd(t *testing.T) {
	for _, testCase := range addTests {
		actual := Add(testCase.input)
		if actual != testCase.output {
			t.Errorf("Oh Snap! Input: %s | Output: %d | Actual: %d", testCase.input, testCase.output, actual)
		}
	}

}
