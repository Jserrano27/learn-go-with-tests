package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrWordNotFound      = DictionaryErr("This word does not exist in the dictionary")
	ErrWordAlreadyExists = DictionaryErr("This word already exists in the dictionary")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
