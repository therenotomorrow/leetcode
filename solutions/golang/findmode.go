package golang

func findMode(root *TreeNode) []int {
	modes := make([]int, 0)
	// took from inordertraversal.go: https://leetcode.com/problems/binary-tree-inorder-traversal/description/
	values := inorderTraversal(root)

	var maxCnt, currCnt, currNum int

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
