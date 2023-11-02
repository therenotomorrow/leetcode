package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func postOrderTraversal(root *structs.TreeNode) ([]int, int) {
	if root == nil {
		return []int{}, 0
	}

	arrL, cntL := postOrderTraversal(root.Left)
	arrR, cntR := postOrderTraversal(root.Right)

	arr := append(arrL, arrR...)
	arr = append(arr, root.Val)

	cnt := 0

	if mathfunc.Sum(arr[0], arr[1:]...)/len(arr) == root.Val {
		cnt++
	}

	return arr, cntL + cntR + cnt
}

func averageOfSubtree(root *structs.TreeNode) int {
	_, cnt := postOrderTraversal(root)
	return cnt
}
