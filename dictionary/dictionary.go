package dictionary

import (
	"errors"
)

// Dictionary stores a map of terms and their definitions
type Dictionary map[string]string

// Search through our dictionary for a term
func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
