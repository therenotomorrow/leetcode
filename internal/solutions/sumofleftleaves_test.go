package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestSumOfLeftLeaves(t *testing.T) {
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
			args: args{root: &structs.TreeNode{
				Val:  3,
				Left: &structs.TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &structs.TreeNode{
					Val:   20,
					Left:  &structs.TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 7, Left: nil, Right: nil},
				},
			}},
			want: 24,
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{Val: 1, Left: nil, Right: nil}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want = %v", got, tt.want)
			}
		})
	}
}
