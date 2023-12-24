package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func minOperations(s string) int {
	cnt := 0 // will count '1'

	for i, digit := range s {
		if i%2 == 0 {
			if digit == '0' {
				cnt++
			}
		} else {
			if digit == '1' {
				cnt++
			}
		}
	}

	// lees correct '1' or '0'
	return mathfunc.Min(cnt, len(s)-cnt)
}
