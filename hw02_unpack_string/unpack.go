package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(a string) (string, error) {
	var str strings.Builder
	b := []rune(a)
	switch {
	case len(b) == 0: // check if len=0
		return "", nil
	case unicode.IsDigit(b[0]): // check if 1st sign is digit
		return "", ErrInvalidString
	}
	// check if 2 digits side by side
	for i := 0; i < len(b)-1; i++ {
		if unicode.IsDigit(b[i]) && unicode.IsDigit(b[i+1]) {
			return "", ErrInvalidString
		}
	}
	for idx, val := range b {
		if !unicode.IsDigit(val) {
			str.WriteString(string(val))
		} else {
			tmpAnsw := str.String()
			str.Reset()
			str.WriteString(tmpAnsw[:len(tmpAnsw)-1])
			strRepStep, _ := strconv.Atoi(string(val))
			str.WriteString(strings.Repeat(string(b[idx-1]), strRepStep))
		}
	}
	return str.String(), nil
}
