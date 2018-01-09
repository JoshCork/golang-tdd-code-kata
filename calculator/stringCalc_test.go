package calculator

import (
	"testing"
)
func TestAdd(t *testing.T) {
	input := "1,2,3"
	expectedValue := 6
	actualValue := Add(input)

	if expectedValue != actualValue {
		t.Errorf("expected: %d, actually got: %d", expectedValue,actualValue)
	}
}