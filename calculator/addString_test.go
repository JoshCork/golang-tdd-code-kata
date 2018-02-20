package calculator

import (
	"testing"
)

var addTests = []struct{
	input		string	// input test case
	ouput 	int			// expected output
	theErr	string	// error message if an error is expected for negative testing
}{
	{"", 0, ""},
	{"1,2,3", 6, ""},
	{"1\n,2,3", 6,""},
	{"\\;\n1\n2;3", 6,""},
	{"-1", 0, "Cannot handle the negativity: ,-1"},
	{"-1,-2,1", 0, "Cannot handle the negativity: ,-1,-2"},

}

func TestAdd(t *testing.T) {
	for _, testCase := range addTests {
		actual,err := Add(testCase.input)
		if err != nil {
			if testCase.theErr != err.Error() {
				t.Errorf("Oh Snap! For Input: %s | Expecting Err: %s | Actual Err: %s", testCase.input, testCase.theErr, err.Error())
			}
		} else if testCase.ouput != actual {
			t.Errorf("Oh Snap! For Input: %s | Expecting Output: %d | Actual Output: %d", testCase.input, testCase.ouput, actual)
		}
	}


}