package dictionary

import ()

// Dictionary stores a map of terms and their definitions
type Dictionary map[string]string

// Search through our dictionary for a term
func (d Dictionary) Search(word string) string {
	return d[word]
}
