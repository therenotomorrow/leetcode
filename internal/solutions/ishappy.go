package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
)

func sumHappy(n int) int {
	newN := 0

	for n > 0 {
		d := n % 10
		newN += d * d
		n /= 10
	}

	return newN
}

func isHappy(n int) bool {
	// find cycle, could use fast and slow runners
	seen := datastruct.NewSet[int]()

	for n > 1 {
		n = sumHappy(n)

		if seen.Contains(n) {
			return false
		}

		seen.Add(n)
	}

	return n == 1
}
