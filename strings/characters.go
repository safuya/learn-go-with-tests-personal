package main

import (
	"fmt"
	"strings"
)

func High(sentence string) string {
	wordHighScore := 0
	highWord := ""

	for _, word := range strings.Fields(sentence) {
		if WordScore(word) > wordHighScore {
			wordHighScore = WordScore(word)
			highWord = word
		}
	}
	return highWord
}

func AlphabetScore() map[string]int {
	a := "abcdefghijklmnopqrstuvwxyz"
	var as = make(map[string]int)

	for i, c := range a {
		as[fmt.Sprintf("%c", c)] = i + 1
	}

	return as
}

func WordScore(word string) int {
	sum := 0
	as := AlphabetScore()

	for _, char := range word {
		sum += as[fmt.Sprintf("%c", char)]
	}

	return sum
}
