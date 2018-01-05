package calculator

import (
	"testing"
)

func TestAdd (t *testing.T) {
	input := ""
	expected := 0
	actual := Add(input)

	if expected != actual {
		t.Errorf("Add returned incorrect result. For input of: %v, expected %v but received output of: %v", input,expected,actual)
	}
}