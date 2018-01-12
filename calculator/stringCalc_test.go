package calculator

import (
	"testing"
)

var addTests = []struct {
	input 		string
	exOutput 	int //expected output
}{
	{"", 0},
	{"1,2,3", 6},
	{"2112", 2112},
	{"1,2,3,4", 10},
}

func TestAdd(t *testing.T) {

	for _, tt := range addTests{
		actualResult := Add(tt.input)
		if actualResult != tt.exOutput {
			t.Errorf("For input of %s the actual result of: %d does not match expected result of: %d",tt.input, actualResult ,tt.exOutput)
		}

	}
}