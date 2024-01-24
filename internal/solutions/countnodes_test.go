package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestCountNodes(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				root: &structs.TreeNode{
					Val: 1,
					Left: &structs.TreeNode{
						Val:   2,
						Left:  &structs.TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 5, Left: nil, Right: nil},
					},
					Right: &structs.TreeNode{
						Val:   3,
						Left:  &structs.TreeNode{Val: 6, Left: nil, Right: nil},
						Right: nil,
					},
				},
			},
			want: 6,
		},
		{name: "smoke 2", args: args{root: nil}, want: 0},
		{name: "smoke 2", args: args{root: &structs.TreeNode{Val: 1, Left: nil, Right: nil}}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want = %v", got, tt.want)
			}
		})
	}
}
