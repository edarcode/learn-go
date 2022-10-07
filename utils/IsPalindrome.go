package utils

import (
	"strings"
)

func IsPalindrome(s string) bool {
	text := strings.ToLower(s)
	lastTextPosition := len(text) - 1
	invertedText := ""

	for position := lastTextPosition; position >= 0; position-- {
		textCharacter := string(text[position])
		invertedText += textCharacter
	}

	return text == invertedText
}
