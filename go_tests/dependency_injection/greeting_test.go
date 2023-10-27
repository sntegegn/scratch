package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Selam")

	actual := buffer.String()
	expected := "Hello Selam"

	if actual != expected {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
