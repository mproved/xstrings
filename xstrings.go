package xstrings

import (
	"unicode"
)

func Substr(text string, from int, length int, dots bool) string {
	runes := []rune(text)
	runesLen := len(runes)

	to := from + length

	if from < 0 {
		from = 0
	}

	if from > runesLen {
		from = runesLen
	}

	if to < 0 {
		to = 0
	}

	if to > runesLen {
		to = runesLen
	}

	result := string(runes[from:to])

	if text != result && dots {
		return result + ".."
	} else {
		return result
	}
}

func IsWhitespace(text string) bool {
	for _, r := range text {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}
