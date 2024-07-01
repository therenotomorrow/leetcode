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

func tree2strBuild(root *TreeNode, builder *strings.Builder) {
	if root == nil {
		return
	}

	builder.WriteString("(" + strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		builder.WriteString(")")

		return
	}

	tree2strBuild(root.Left, builder)

	if root.Left == nil {
		// save one-to-one mapping
		builder.WriteString("()")
	}

	tree2strBuild(root.Right, builder)

	builder.WriteString(")")
}
