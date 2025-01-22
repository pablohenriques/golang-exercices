package sequential

import (
	"fmt"
	"testing"
)

func MessageNumber(number int) string {
	return fmt.Sprintf("Number is %d", number)
}

func TestMessageNumber(t *testing.T) {
	result := MessageNumber(1)
	expected := "Number is 1"

	if result != expected {
		t.Errorf("Returned message is different")
	}
}
