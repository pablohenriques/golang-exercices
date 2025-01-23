package sequential

import (
	"testing"
)

func SumTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func TestSumTwoNumbers(t *testing.T) {
	result := SumTwoNumbers(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Values are differents")
	}
}
