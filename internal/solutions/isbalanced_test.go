package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{root: &structs.TreeNode{
				Val:  3,
				Left: &structs.TreeNode{Val: 9},
				Right: &structs.TreeNode{
					Val:   20,
					Left:  &structs.TreeNode{Val: 15},
					Right: &structs.TreeNode{Val: 7},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{
				Val: 1,
				Left: &structs.TreeNode{
					Val: 2,
					Left: &structs.TreeNode{
						Val:   3,
						Left:  &structs.TreeNode{Val: 4},
						Right: &structs.TreeNode{Val: 4},
					},
					Right: &structs.TreeNode{Val: 3},
				},
				Right: &structs.TreeNode{Val: 2},
			}},
			want: false,
		},
		{
			name: "smoke 3",
			args: args{root: nil},
			want: true,
		},
		{
			name: "test 203: wrong answer",
			args: args{root: &structs.TreeNode{
				Val: 1,
				Left: &structs.TreeNode{
					Val: 2,
					Left: &structs.TreeNode{
						Val:  3,
						Left: &structs.TreeNode{Val: 4},
					},
				},
				Right: &structs.TreeNode{
					Val: 2,
					Right: &structs.TreeNode{
						Val:   3,
						Right: &structs.TreeNode{Val: 4},
					},
				},
			}},
			want: false,
		},
		{
			name: "test 34: wrong answer",
			args: args{root: &structs.TreeNode{
				Val: 1,
				Left: &structs.TreeNode{
					Val: 2,
					Left: &structs.TreeNode{
						Val:   3,
						Left:  &structs.TreeNode{Val: 4},
						Right: &structs.TreeNode{Val: 4},
					},
					Right: &structs.TreeNode{Val: 3},
				},
				Right: &structs.TreeNode{Val: 2},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want = %v", got, tt.want)
			}
		})
	}
}
