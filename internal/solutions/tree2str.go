package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"strconv"
	"strings"
)

func tree2str(root *structs.TreeNode) string {
	var sb strings.Builder

	tree2strBuild(root, &sb)

	return sb.String()[1 : sb.Len()-1]
}

func tree2strBuild(root *structs.TreeNode, sb *strings.Builder) {
	if root == nil {
		return
	}

	sb.WriteString("(" + strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		sb.WriteString(")")
		return
	}

	tree2strBuild(root.Left, sb)

	if root.Left == nil {
		// save one-to-one mapping
		sb.WriteString("()")
	}

	tree2strBuild(root.Right, sb)

	sb.WriteString(")")
}
