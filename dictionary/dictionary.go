package dictionary

import (
	"errors"
)

// Dictionary stores a map of terms and their definitions
type Dictionary map[string]string

var (
	// ErrNotFound is thrown when we cannot find the definition
	ErrNotFound = errors.New("could not find the word you were looking for")

	// ErrWordExists is thrown when attempting to add a word that's already defined
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search through our dictionary for a term
func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Remove and entry from the dictionary
func (d Dictionary) Remove(word string) {
	delete(d, word)
}
