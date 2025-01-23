package sequential

import "testing"

func ConverterMeasures(meters int) int {
	return meters * 100
}

func TestConverterMeasures(t *testing.T) {
	var result int = ConverterMeasures(1)
	var expected int = 100

	if result != expected {
		t.Errorf("Result is different")
	}
}
