package solutions

import (
	"fmt"
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"math/rand"
	"strconv"
)

func findDifferentBinaryString(nums []string) string {
	uniq := datastruct.NewSet[int]()

	for _, num := range nums {
		val, _ := strconv.ParseInt(num, 2, 32)
		uniq.Add(int(val))
	}

	found := -1
	for {
		found = rand.Intn(len(nums) + 1)
		if !uniq.Contains(found) {
			break
		}
	}

	return fmt.Sprintf("%0*s", len(nums), strconv.FormatInt(int64(found), 2))
}
