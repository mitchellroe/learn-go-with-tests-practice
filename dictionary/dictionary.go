package dictionary

import (
	"errors"
)

// Dictionary stores a map of terms and their definitions
type Dictionary map[string]string

// ErrNotFound is thrown when we cannot find the definition
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search through our dictionary for a term
func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}
