package dictionary

type Dictionary map[string]string

type DictionaryError string

// var dictionary = map[string]string{}

// // OR

// var dictionary = make(map[string]string)

// Never have an empty map
const (
	ErrNoWordFound      = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("could not add a new word, word exist!")
	ErrUpdateWordExists = DictionaryError("could now update a new world, try add")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {

		return "", ErrNoWordFound
	}

	return definition, nil
}

func (d Dictionary) AddWord(word, definition string) error {
	_, err := d.Search(word)
	if err != ErrNoWordFound {
		return ErrWordExists
	}
	d[word] = definition
	return nil
	// switch err {
	// case ErrNoWordFound:
	// 	d[word] = definition
	// case nil:
	// 	return ErrWordExists
	// default:
	// 	return err
	// }
	// return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	// we can also use switch statement.
	if err != ErrNoWordFound {
		d[word] = newDefinition
		return nil
	}
	return ErrUpdateWordExists
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
