package common

import (
	"bytes"
	"unicode"
)

// DateFormat is a format string used tim time.Time.Format()
const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func CamelToSnake(s string) string {
	buf := bytes.Buffer{}

	isPrevUpper := false
	lastRune := ' '

	for _, v := range s {
		if unicode.IsUpper(v) {
			if !isPrevUpper && lastRune != ' ' {
				buf.WriteRune('_')
			}
			isPrevUpper = true
		} else {
			isPrevUpper = false
		}

		buf.WriteRune(unicode.ToLower(v))
		lastRune = v
	}

	return buf.String()
}
