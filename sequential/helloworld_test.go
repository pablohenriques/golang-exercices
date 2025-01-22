package sequential

import (
	"testing"
)

func HelloWorld() string {
	return "Hello World"
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld()

	if result != "Hello World" {
		t.Errorf("Hello World() function is not correct.")
	}
}
