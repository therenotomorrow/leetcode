package golang

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	const maxVal int = 1e18 - 1

	newPalindrome := func(num int) (int, error) {
		runes := []rune(strconv.Itoa(num))

		// just fill first part with the same mirror variables
		for l, r := 0, len(runes)-1; l <= r; l, r = l+1, r-1 {
			runes[r] = runes[l]
		}

		return strconv.Atoi(string(runes))
	}

	var (
		num, _            = strconv.Atoi(n)
		curr, left, right int
	)

	// ---- previous value by binary search
	left, right = 0, num-1
	prev := math.MinInt

	for left <= right {
		mid := (right-left)/Half + left
		curr, _ = newPalindrome(mid)

		if curr < num {
			prev = curr
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// ---- next value by binary search
	left, right = num+1, maxVal
	next := math.MinInt

	for left <= right {
		mid := (right-left)/Half + left
		curr, _ = newPalindrome(mid)

		if curr > num {
			next = curr
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if Abs(prev-num) <= Abs(next-num) {
		return strconv.Itoa(prev)
	}

	return strconv.Itoa(next)
}
