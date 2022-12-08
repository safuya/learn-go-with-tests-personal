package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sophie", "")
		want := "Hello, Sophie."
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Nobody", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world."
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sofia", "Spanish")
		want := "Hola, Sofia."
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sophie", "French")
		want := "Bonjour, Sophie."
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
