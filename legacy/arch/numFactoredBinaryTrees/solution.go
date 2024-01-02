package numFactoredBinaryTrees

import (
	"sort"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)

	binaryTrees := make(map[int]int, len(arr))
	// each node is a tree without leafs
	for _, node := range arr {
		binaryTrees[node] = 1
	}

	for _, root := range arr {
		for _, node := range arr {
			if root == node {
				break
			}

			if root%node != 0 {
				continue
			}

			val, ok := binaryTrees[root/node]
			if !ok {
				continue
			}

			binaryTrees[root] += val * binaryTrees[node]
		}
	}

	sum := 0

	for _, cntTrees := range binaryTrees {
		sum += cntTrees
	}

	return sum % structs.MOD
}
