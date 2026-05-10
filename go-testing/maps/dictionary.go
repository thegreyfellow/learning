package main

const (
	ErrNotFound         = DictionaryErr("keyword not found")
	ErrWordExists       = DictionaryErr("keyword already exists")
	ErrWordDoesNotExist = DictionaryErr("keyword does not exist")
)

type DictionaryErr string

type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(keyword string) (string, error) {
	value, ok := d[keyword]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}
