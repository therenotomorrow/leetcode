package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func inOrderTraversal(root *structs.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	arr := inOrderTraversal(root.Left)
	arr = append(arr, root.Val)
	arr = append(arr, inOrderTraversal(root.Right)...)

	return arr
}

func findMode(root *structs.TreeNode) []int {
	modes := make([]int, 0)
	values := inOrderTraversal(root)

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
