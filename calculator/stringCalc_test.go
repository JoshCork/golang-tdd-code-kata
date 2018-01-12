package calculator

import (
	"testing"
)

var addTests = []struct {
	input 		string
	expected 	int
}{
	{"",0},
	{"1,2,3", 6},
	{"2112", 2112},
}

func TestAdd(t *testing.T) {
	// some table driven tests

	for _,tt := range addTests {
		actual := Add(tt.input)
		if actual != tt.expected {
			t.Errorf("ERROR: input: %s | expected result: %d | actual result: %d",tt.input, tt.expected, actual)
		}
	}

}