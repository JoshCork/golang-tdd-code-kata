package calculator
import (
	"testing"
)

var addTests = []struct {
	input		string		// input test case
	output	int				// expected output
	theErr	string		// if expecting an error this is the expected message
}{
	{"", 0,""},
	{"1,2,3", 6,""},
	{"1\n2,3", 6, ""},
	{"\\;\n1\n2;3", 6, ""},
	{"-1", -1, "Cannot handle the negativity ,-1"},
	{"-1,-2,-3", -1, "Cannot handle the negativity ,-1,-2,-3"},
	{"2112", 0, ""},
	{"2,1001", 2, ""},
	{"\\[***]\n1\n2***3", 6, ""},
	{"\\[***][;][||]\n1***2;3||4",10,""},
}

func TestAdd(t *testing.T) {
	for _,testCase := range addTests {
		actual,err := Add(testCase.input)
		if err != nil {
			if testCase.theErr != err.Error() {
				t.Errorf("Oh Snap! For Input: %s | Expecting Error: %s | Actual Error: %s", testCase.input, testCase.theErr, err.Error())
			}
		} else if actual != testCase.output {
			t.Errorf("Oh Snap! For input: %s | Expected Output: %d | Actual: %d", testCase.input, testCase.output, actual)
		}
	}

}