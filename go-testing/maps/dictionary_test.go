package main

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("search for existing keyword", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("search for none existing keyword", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("add new keyword", func(t *testing.T) {
		word := "funny"
		definition := "someone or something that makes you laugh"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("add existing keyword", func(t *testing.T) {
		word := "test"
		want, _ := dictionary.Search(word)
		definition := "test definition but version 2.0"
		err := dictionary.Add(word, definition)

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is an UPDATED DEFINITION"
		dictionary := Dictionary{word: definition}
		dictionary.Update(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update non-existing word", func(t *testing.T) {
		word := "test"
		definition := "this is an UPDATED DEFINITION"
		dictionary := Dictionary{}
		got := dictionary.Update(word, definition)

		assertErrors(t, got, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		definition := "some test definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertErrors(t, err, nil)

		_, err2 := dictionary.Search(word)
		assertErrors(t, err2, ErrNotFound)
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}
