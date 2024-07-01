package golang

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func findDifferentBinaryString(nums []string) string {
	uniq := NewSet[int]()

	for _, num := range nums {
		val, _ := strconv.ParseInt(num, 2, 32)
		uniq.Add(int(val))
	}

	var found int

	for {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(nums)+1)))

		found = int(idx.Int64())
		if !uniq.Contains(found) {
			break
		}
	}

	return fmt.Sprintf("%0*s", len(nums), strconv.FormatInt(int64(found), 2))
}
