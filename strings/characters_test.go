package main

import "testing"

func TestHigh(t *testing.T) {
	got := High("man i need a taxi up to ubud")
	want := "taxi"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestAlphabetScore(t *testing.T) {
	as := AlphabetScore()
	got := as["z"]
	want := 26

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestWordScore(t *testing.T) {
	got := WordScore("abc")
	want := 6

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
