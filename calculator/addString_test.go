package calculator

import "testing"

var addTests = []struct {
	input  string //test case input
	output int    //expected output from the test
	theErr string //error message if expecting an error output
}{
	{"", 0, ""},
	{"2112", 0, ""},
	{"1,2,3", 6, ""},
	{"1\n2,3", 6, ""},
	{"\\;\n1\n2;3", 6, ""},
	{"-1", -1, "Cannot handle the negativity,-1"},
	{"-1,1,2,-3", -1, "Cannot handle the negativity,-1,-3"},
	{"2112,1,2", 3, ""},
}

func TestAdd(t *testing.T) {

	for _, testCase := range addTests {
		actual, err := Add(testCase.input)
		if err != nil {
			// there was an error, was it what we were expecting?
			if err.Error() != testCase.theErr {
				t.Errorf("Oh Snap! For Input: %s | Expecting Error: %s | Actual Error: %s", testCase.input, testCase.theErr, err.Error())
			}
		} else {
			if actual != testCase.output {
				t.Errorf("Oh Snap! For Input: %s | Expecting Output: %d | Actual Output: %d", testCase.input, testCase.output, actual)
			}
		}
	}

}
