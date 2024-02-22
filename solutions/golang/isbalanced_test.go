package golang

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 2, Left: nil, Right: nil},
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
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
				},
			}},
			want: false,
		},
		{
			name: "test 34: wrong answer",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 2, Left: nil, Right: nil},
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
