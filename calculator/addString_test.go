package calculator

import (
	"testing"
)

var AddTests = []struct {
	input  string // test input
	output int    // expected output
}{
	{"", 0},
	{"2112", 2112},
	{"1,2,3", 6},
	{"1\n2,3", 6},
	{"\\;\n1;2;3", 6},
}

func TestAdd(t *testing.T) {

	for _, testCase := range AddTests {
		actual := Add(testCase.input)
		if actual != testCase.output {
			t.Errorf("Oh Snap! Input: %s | Expected Output: %d | Actual Output: %d", testCase.input, testCase.output, actual)
		}
	}
}
