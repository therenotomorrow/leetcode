package golang

import "math"

func minWindow(s string, t string) string {
	var (
		cntT                  = make(map[byte]int)
		cntS                  = make(map[byte]int)
		minLen                = math.MaxInt
		foundL, foundR, match int
	)

	for i := range t {
		cntT[t[i]]++
	}

	match = len(cntT)

	for left, right := 0, 0; right < len(s); right++ {
		cntS[s[right]]++

		// because we go from left to right - each match will cost
		if cntS[s[right]] == cntT[s[right]] {
			match--
		}

		for ; match == 0; left++ {
			if right-left+1 < minLen {
				minLen = right - left + 1
				foundR = right + 1
				foundL = left
			}

			unused := s[left]

			if cntS[unused] > cntT[unused] {
				cntS[unused]--

				continue
			}

			if cntS[unused] < cntT[unused] {
				match++

				continue
			}

			break
		}
	}

	if minLen == math.MaxInt {
		return ""
	}

	return s[foundL:foundR]
}
