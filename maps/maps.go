package maps

//import "errors"

const (
    ErrNotFound = DictionaryErr("Could not find the word you were looking for")
    ErrWordExists = DictionaryErr("Word already exist in dictionary")
    ErrWordDoesNotExist = DictionaryErr("Cannot update word cause it does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

func (dict Dictionary) Search(word string) (string, error) {
    definition, ok := dict[word]
    if !ok {
        return "", ErrNotFound
    }
    return definition, nil
}
func (dict Dictionary) Add(word string, definition string) error {
    _, err := dict.Search(word)

    switch err {
    case ErrNotFound:
        dict[word] = definition
    case nil:
        return ErrWordExists
    default:
        return err
    }
    return nil
}
func (dict Dictionary) Update(word,  definition string) error {
    _, err := dict.Search(word)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        dict[word] = definition
    default:
        return err
    }
    return nil
}
func (dict Dictionary) Delete(word string) {
    delete(dict, word)
}
