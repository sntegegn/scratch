package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("Selam")
	expected := "Hello Selam"

	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
}
