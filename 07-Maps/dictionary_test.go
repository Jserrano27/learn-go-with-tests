package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, expected, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		assertError(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add("test", "Another definition")

		assertError(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is the original definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is the new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("a new word", func(t *testing.T) {
		word := "test"
		definition := "this is the original definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a sample definition"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertError(t, err, ErrWordNotFound)
}

func assertStrings(t *testing.T, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("Expected was %q, but got %q instead", expected, got)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Error finding de added word:", err)
	}

	if got != definition {
		t.Errorf("Expected was %q, but got %q instead", definition, got)
	}
}

func assertError(t *testing.T, got, expected error) {
	t.Helper()

	if expected != got {
		t.Errorf("Expected error was %q, but got %q instead", expected, got)
	}

	if got == nil {
		if expected == nil {
			return
		}
		t.Fatal("Expected an error but got none")
	}
}
