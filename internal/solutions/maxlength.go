package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func maxLength(arr []string) int {
	var dynamic func(curr int, str string) int

	dynamic = func(currIdx int, currStr string) int {
		set := datastruct.NewSet[rune]()
		for _, r := range currStr {
			set.Add(r)
		}

		maxLen := set.Size()

		if maxLen != len(currStr) {
			return 0
		}

		for i := currIdx; i < len(arr); i++ {
			maxLen = mathfunc.Max(maxLen, dynamic(i+1, currStr+arr[i]))
		}

		return maxLen
	}

	return dynamic(0, "")
}
