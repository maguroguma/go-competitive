package runestring

import (
	"strings"
)

// UpperRune is rune version of `strings.ToUpper()`.
func UpperRune(r rune) rune {
	str := strings.ToUpper(string(r))
	return []rune(str)[0]
}

// LowerRune is rune version of `strings.ToLower()`.
func LowerRune(r rune) rune {
	str := strings.ToLower(string(r))
	return []rune(str)[0]
}

// ToggleRune returns a upper case if an input is a lower case, v.v.
func ToggleRune(r rune) rune {
	var str string
	if 'a' <= r && r <= 'z' {
		str = strings.ToUpper(string(r))
	} else if 'A' <= r && r <= 'Z' {
		str = strings.ToLower(string(r))
	} else {
		str = string(r)
	}
	return []rune(str)[0]
}

// ToggleString iteratively calls ToggleRune, and returns the toggled string.
func ToggleString(s string) string {
	inputRunes := []rune(s)
	outputRunes := make([]rune, 0, len(inputRunes))
	for _, r := range inputRunes {
		outputRunes = append(outputRunes, ToggleRune(r))
	}
	return string(outputRunes)
}
