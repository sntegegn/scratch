package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("testing a word in the dictionary", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertString(t, actual, expected)
	})
	t.Run("testing a word not in the dictionary", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("Expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("hello", "greetings")
		assertError(t, err, nil)
		assertDefination(t, dictionary, "hello", "greetings")
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"hello": "greeting"}
		err := dictionary.Add("hello", "another greeting")

		assertError(t, err, ErrWordExists)
	})

}

func assertDefination(t *testing.T, dictionary Dictionary, word, defination string) {
	actual, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertString(t, actual, defination)
}

func assertString(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
