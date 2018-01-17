package calculator

import (
	"testing"
)

var addTests = []struct{
	inputV 		string
	expectedV int
}{
	{"",0},
	{"1,2,3",6},
	{"1,2,3,4",10},
	{"2112",2112},
	{"1\n2,3",6},
}

func TestAdd(t *testing.T) {
	for _, testcase := range addTests {
		actual := Add(testcase.inputV)
		if actual != testcase.expectedV {
			t.Errorf("\nERROR! - input: %s | expected output: %d | actual output: %d\n", testcase.inputV, testcase.expectedV, actual)
		}
	}
}