package sequential

import "testing"

func GradePointAverage(n1, n2, n3, n4 int) int {
	return (n1 + n2 + n3 + n4) / 4
}

func TestGradePointAverage(t *testing.T) {
	result := GradePointAverage(7, 7, 7, 7)
	expected := 7

	if result != expected {
		t.Errorf("Average is not correct")
	}
}
