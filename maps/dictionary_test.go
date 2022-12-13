package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		AssertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		AssertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, _ := dictionary.Search("test")

		AssertString(t, got, want)
	})

	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "I am immutable")
		definition, _ := dictionary.Search("test")

		AssertError(t, err, ErrWordExists)
		AssertString(t, definition, "this is just a test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "old definition"}

		dictionary.Update("test", "new definition")
		definition, _ := dictionary.Search("test")

		AssertString(t, definition, "new definition")
	})

	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Update("test", "I shouldn't exist")
		_, err := dictionary.Search("test")

		AssertError(t, err, ErrNotFound)
	})
}

func AssertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("Wanted error %q but didn't get an error", want)
	}

	if got != want {
		t.Errorf("Got error %q wanted error %q", got, want)
	}
}
