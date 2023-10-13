package sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return createNode(nums, 0, len(nums)-1)
}

func createNode(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (right + left) / 2

	node := &TreeNode{Val: nums[mid]}
	node.Left = createNode(nums, left, mid-1)
	node.Right = createNode(nums, mid+1, right)

	return node
}
