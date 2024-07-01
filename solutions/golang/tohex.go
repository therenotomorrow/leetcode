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

	const (
		AHex = Digits + iota
		BHex
		CHex
		DHex
		EHex
		FHex
		Hex
	)

	stack := NewStack[rune]()

	for ; num > 0; num /= Hex {
		switch val := num % Hex; val {
		case AHex, BHex, CHex, DHex, EHex, FHex:
			stack.Push(rune(alphaPrefix + val))
		default:
			stack.Push(rune(digitPrefix + val))
		}
	}

	var builder strings.Builder

	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		builder.WriteRune(val)
	}

	return builder.String()
}
