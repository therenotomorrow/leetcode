package golang

import (
	"strconv"
	"strings"
)

func tree2str(root *TreeNode) string {
	var sb strings.Builder

	tree2strBuild(root, &sb)

	return sb.String()[1 : sb.Len()-1]
}

func tree2strBuild(root *TreeNode, sb *strings.Builder) {
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
