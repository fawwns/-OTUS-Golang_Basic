package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	var escape bool
	var prev rune
	var hasPrev bool

	for _, r := range str {
		if escape {
			if r != '\\' && !unicode.IsDigit(r) {
				return "", ErrInvalidString
			}
			prev = r
			result.WriteRune(r)
			hasPrev = true
			escape = false
			continue
		}

		if r == '\\' {
			escape = true
			continue
		}

		if unicode.IsDigit(r) {
			if !hasPrev {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(r))
			if count == 0 {
				// удаляем предыдущий символ
				strSoFar := result.String()
				result.Reset()
				result.WriteString(strSoFar[:len(strSoFar)-1])
				hasPrev = false
				prev = 0
				continue
			}
			// удаляем уже добавленный один раз символ
			strSoFar := result.String()
			result.Reset()
			result.WriteString(strSoFar[:len(strSoFar)-1])
			result.WriteString(strings.Repeat(string(prev), count))
			hasPrev = false
			prev = 0
			continue
		}

		result.WriteRune(r)
		prev = r
		hasPrev = true
	}

	if escape {
		return "", ErrInvalidString
	}

	return result.String(), nil
}
