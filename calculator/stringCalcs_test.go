package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	inputValue := ""
	expectedValue := 0
	myAddition := Add(inputValue)

	if myAddition != expectedValue {
		t.Errorf("myAddtion is expecting %d, but got %d", expectedValue,myAddition)
	}
}