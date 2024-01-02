package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func maximumAverageSubtree(root *structs.TreeNode) float64 {
	_, _, avg := inOrderCounter(root)

	return avg
}

func inOrderCounter(root *structs.TreeNode) (int, int, float64) {
	if root == nil {
		return 0, 0, 0
	}

	sum := 0
	cnt := 0

	// ---- left
	sumL, cntL, avgL := inOrderCounter(root.Left)

	sum += sumL
	cnt += cntL

	// ---- right
	sumR, cntR, avgR := inOrderCounter(root.Right)

	sum += sumR
	cnt += cntR

	// ---- postorder
	sum += root.Val
	cnt++

	return sum, cnt, mathfunc.Max(float64(sum)/float64(cnt), avgL, avgR)
}
