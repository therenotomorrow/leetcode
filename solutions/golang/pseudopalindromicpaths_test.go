package golang

import (
	"testing"
)

func TestPseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
				},
			},
			want: 2,
		},
		{
			name: "smoke 2",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: &TreeNode{Val: 1, Left: nil, Right: nil},
						},
					},
					Right: &TreeNode{Val: 1, Left: nil, Right: nil},
				},
			},
			want: 1,
		},
		{
			name: "smoke 3",
			args: args{root: &TreeNode{Val: 9, Left: nil, Right: nil}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want = %v", got, tt.want)
			}
		})
	}
}
