package golang

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
				subRoot: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
				subRoot: nil,
			},
			want: false,
		},
		{
			name: "test 162: wrong answer",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: nil,
				},
				subRoot: &TreeNode{Val: 1, Left: nil, Right: nil},
			},
			want: true,
		},
		{
			name: "test 177: wrong answer",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: nil,
					},
				},
				subRoot: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
