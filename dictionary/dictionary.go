package dictionary

import (
	"errors"
)

// Dictionary stores a map of terms and their definitions
type Dictionary map[string]string

// ErrNotFound is thrown when we cannot find the definition
var ErrNotFound = errors.New("could not find the word you were looking for")

// ErrWordExists is thrown when attempting to add a word that's already defined
var ErrWordExists = errors.New("word already has a definition")

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
	if d[word] != "" {
		return ErrDuplicate
	}
	d[word] = definition
	return nil
}

// Remove and entry from the dictionary
func (d Dictionary) Remove(word string) {
	delete(d, word)
}
