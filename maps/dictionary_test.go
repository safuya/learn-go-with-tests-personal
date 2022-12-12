package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		AssertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}
		AssertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, _ := dictionary.Search("test")

		AssertString(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "I am immutable")
		definition, _ := dictionary.Search("test")

		AssertError(t, err, ErrWordExists)
		AssertString(t, definition, "this is just a test")
	})
}

func AssertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q wanted error %q", got, want)
	}
}
