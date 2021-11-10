package dictionary

import (
	"errors"
)

// Making alias of types, customized types.
// Dictionary type.
type Dictionary map[string]string

var (
	errNoneFound = errors.New("The word is not found.")
	errWordExist = errors.New("The word is already in the dictionary.")
)

// Search for a word in the dictionary.
func (d Dictionary) SearchWord(word string) (string, error) {
	definition, exist := d[word]
	if exist {
		return definition, nil
	}
	return "", errNoneFound
}

// Add a word in the dictionary.
func (d Dictionary) Add(word string, def string) error {
	_, err := d.SearchWord(word)
	if err == errNoneFound {
		d[word] = def
		return nil
	}
	return errWordExist
}

// Update the definition of the word.
func (d Dictionary) Update(word string, def string) error {
	_, err := d.SearchWord(word)
	if err == nil {
		d[word] = def
		return nil
	}
	return errNoneFound
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}