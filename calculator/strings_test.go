package calculator

import "testing"

func TestAdd(t *testing.T) {
	input1 := ""
	total := Add(input1)
	if total != 0 {
		t.Errorf("Addition of the two strings was incorrect, got: %d, expecting: %d.",total, 0)
	}
}