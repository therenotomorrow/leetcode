package golang

import (
	"math"
	"strings"
)

const (
	alphaPrefix = 87
	digitPrefix = 48
)

func toHex(num int) string {
	if num < 0 {
		// two’s complement method
		num = math.MaxUint32 + (num + 1)
	}

	if num == 0 {
		// edge case
		return "0"
	}

	stack := NewStack[rune]()

	for ; num > 0; num /= 16 {
		switch val := num % 16; val {
		case 10, 11, 12, 13, 14, 15: //nolint:gomnd
			stack.Push(rune(alphaPrefix + val))
		default:
			stack.Push(rune(digitPrefix + val))
		}
	}

	var sb strings.Builder

	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		sb.WriteRune(val)
	}

	return sb.String()
}
