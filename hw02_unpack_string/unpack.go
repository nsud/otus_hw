package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	var buildS strings.Builder
	var s = []rune(input)
	firstSymbol := s[0]
	lastSymbol := s[len(s)-1]
	if unicode.IsDigit(firstSymbol) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(s)-1; i++ {
		switch {
		case (unicode.IsSymbol(s[i]) || unicode.IsLetter(s[i])) && !unicode.IsDigit(s[i+1]):
			buildS.WriteString(string(s[i]))
		case (unicode.IsSymbol(s[i]) || unicode.IsLetter(s[i])) && unicode.IsDigit(s[i+1]):
			if n, err := strconv.Atoi(string(s[i+1])); err == nil {
				temp := strings.Repeat(string(s[i]), n)
				buildS.WriteString(temp)
			}
		case unicode.IsDigit(s[i]) && unicode.IsDigit(s[i+1]):
			return "", ErrInvalidString
		}
	}
	if !unicode.IsDigit(lastSymbol) {
		buildS.WriteString(string(lastSymbol))
	}

	return buildS.String(), nil
}
