package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Mary", "English")
		expected := "Hello Mary"

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("saying 'Hello world' when empty string is provided", func(t *testing.T) {
		actual := Hello("", "English")
		expected := "Hello world"

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("John", "Spanish")
		expected := "Hola John"

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Hello in French", func(t *testing.T) {
		actual := Hello("Bob", "French")
		expected := "Bonjour Bob"

		assertCorrectMessage(t, expected, actual)
	})
}

// testing.T and testing.B are structs whose pointer recivers implements the tetsting.TB interface
func assertCorrectMessage(t *testing.T, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
}
