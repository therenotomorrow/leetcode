package myAtoi

import (
	"strings"
	"unicode"
)

func TrimSpaceLeft(s string) string {
	for i, r := range s {
		if !unicode.IsSpace(r) {
			return s[i:]
		}
	}

	return s
}

func ParseValidNumber(s string) string {
	var sb strings.Builder

	switch {
	case s == "":
		s = "+"
	case s[0] == '+', s[0] == '-':
		sb.WriteByte(s[0])
		s = s[1:]
	default:
		sb.WriteByte('+')
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			break
		}
		sb.WriteRune(r)
	}

	return sb.String()
}

func ParseInt32(s string) int {
	var num int64

	isNeg := s[0] == '-'

	for _, r := range s[1:] {
		num *= 10
		num += int64(r - '0')

		if isNeg {
			if num > 2147483648 {
				return -2147483648
			}
		} else {
			if num > 2147483647 {
				return 2147483647
			}
		}
	}

	if isNeg {
		num = -num
	}

	return int(num)
}

func myAtoi(s string) int {
	s = ParseValidNumber(TrimSpaceLeft(s))

	if s == "" {
		return 0
	}

	return ParseInt32(s)
}
