package maps

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("couldn't find word in the dictionary")
	ErrWordExists = errors.New("word already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

}
