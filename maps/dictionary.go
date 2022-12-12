package dictionary

import "errors"

var ErrNotFound = errors.New("the word you are looking for can't be found")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
