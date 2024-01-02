package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var count int
	calc(root, make(map[int]int), 0, targetSum, &count)
	return count
}

func calc(node *TreeNode, hSet map[int]int, currSum int, targetSum int, count *int) {
	if node == nil {
		return
	}

	currSum += node.Val
	// add if start from the root node
	if currSum == targetSum {
		*count += 1
	}

	// if yes - we will add subtree that starts not from root node - tricky
	*count += hSet[currSum-targetSum]

	hSet[currSum] += 1

	calc(node.Left, hSet, currSum, targetSum, count)
	calc(node.Right, hSet, currSum, targetSum, count)

	hSet[currSum] -= 1
}
