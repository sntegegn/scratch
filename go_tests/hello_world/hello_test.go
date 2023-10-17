package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Selam")
		expected := "Hello Selam!"

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("saying 'Hello world' when empty string is provided", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello world"

		assertCorrectMessage(t, expected, actual)
	})
}

// testing.T and testing.B are structs whose pointer recivers implements the tetsting.TB interface
func assertCorrectMessage(t testing.TB, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
}
