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

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		assertDefination(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	dictionary.Delete("test")

	_, err := dictionary.Search("test")
	assertError(t, err, ErrNotFound)
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
