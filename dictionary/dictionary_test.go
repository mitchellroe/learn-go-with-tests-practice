package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("add a new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add a word already defined", func(t *testing.T) {
		word := "foo"
		definition := "bar"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		err2 := dictionary.Add(word, definition)
		assertError(t, err2, ErrDuplicate)
	})

}

func TestRemove(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	dictionary.Remove("test")
	_, err := dictionary.Search("test")
	assertError(t, err, ErrNotFound)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	want := definition
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, want)
}
