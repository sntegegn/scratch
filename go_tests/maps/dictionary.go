package maps

const (
	ErrNotFound         = DictionaryErr("couldn't find word in the dictionary")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, definition string) error {
	if _, err := d.Search(word); err != nil {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
