package calculator

import "testing"

var addTests = []struct{
	inputValue			string
	expectedValue 	int

}{
	{"",0},
	{"1,2,3", 6},
	{"2112",2112},
	{"1\n2,3", 6},
	{"//;\n1;2",3},
}

func TestAdd(t *testing.T) {

for _,testCase := range addTests {
	actual := Add(testCase.inputValue)

	if actual != testCase.expectedValue {
		t.Errorf("ERROR: input: %s | expected: %d | actual: %d" , testCase.inputValue, testCase.expectedValue, actual)
	}
}

}