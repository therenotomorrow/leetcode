package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func findMode(root *structs.TreeNode) []int {
	modes := make([]int, 0)
	// took from inordertraversal.go: https://leetcode.com/problems/binary-tree-inorder-traversal/description/
	values := inorderTraversal(root)

	maxCnt := 0
	currCnt := 0
	currNum := 0

	for _, num := range values {
		if num == currNum {
			currCnt++
		} else {
			currNum = num
			currCnt = 1
		}

		// we found the biggest one
		if currCnt > maxCnt {
			modes = make([]int, 0)
			maxCnt = currCnt
		}

		if currCnt == maxCnt {
			modes = append(modes, currNum)
		}
	}

	return modes
}
