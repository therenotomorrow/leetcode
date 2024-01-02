package getLengthOfOptimalCompression

import (
	"cmp"
	"math"
)

func getLengthOfOptimalCompression(s string, k int) int {
	var (
		c       = make(map[int]int)
		dynamic func(curr int, last byte, cnt int, k int) (minLen int)
	)

	dynamic = func(curr int, last byte, cnt int, k int) (minLen int) {
		key := curr*101*27*101 + int(last-'a')*101*101 + cnt*101 + k

		if val, ok := c[key]; ok {
			return val
		}

		if k < 0 {
			return math.MaxInt
		}

		if curr == len(s) {
			return 0
		}

		minLenAfterDel := dynamic(curr+1, last, cnt, k-1)
		minLenAfterKeep := 0

		if lasCurr := s[curr]; lasCurr == last {
			one := 0
			if cnt == 1 || cnt == 9 || cnt == 99 {
				one = 1
			} else {
				one = 0
			}

			minLenAfterKeep = dynamic(curr+1, last, cnt+1, k) + one
		} else {
			minLenAfterKeep = dynamic(curr+1, lasCurr, 1, k) + 1
		}

		minLen = Min(minLenAfterDel, minLenAfterKeep)

		c[key] = minLen

		return minLen
	}

	return dynamic(0, '0', 0, k)
}

func Min[T cmp.Ordered](nums ...T) T {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}
