package dictionary

import "testing"

func TestDictionary(t *testing.T) {

	t.Run("known dictionary word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a simple test"}

		got, _ := dictionary.Search("test")
		want := "This is a simple test"

		assertString(t, got, want)
	})

	t.Run("unknown dictionary word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a simple test"}

		_, err := dictionary.Search("tester")
		want := "could not find the word you were looking for"

		assertError(t, err, ErrNoWordFound)

		assertString(t, err.Error(), want)
	})
}

func TestAddWord(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is a simple test"
		err := dictionary.AddWord(word, definition)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "This is a simple test"
		dictionary := Dictionary{word: definition}
		err := dictionary.AddWord(word, definition)

		assertError(t, err, ErrWordExists)

		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test when we update existing word", func(t *testing.T) {
		word := "test"
		definition := "This is a simple test"
		newDefinition := "updated simple test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("test when we update on new word", func(t *testing.T) {
		word := "test"
		definition := "This is a simple test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrUpdateWordExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "This is a simple test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	if err != ErrNoWordFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	want := definition

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
