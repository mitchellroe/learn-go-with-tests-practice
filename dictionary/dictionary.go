package main

import ()

// Search through our dictionary for a term
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
