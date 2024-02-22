package golang

import (
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
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
			args: args{root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			}},
			want: 24,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{Val: 1, Left: nil, Right: nil}},
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
