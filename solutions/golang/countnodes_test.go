package golang

import (
	"testing"
)

func TestCountNodes(t *testing.T) {
	t.Parallel()

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
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{Val: 5, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: nil,
					},
				},
			},
			want: 6,
		},
		{name: "smoke 2", args: args{root: nil}, want: 0},
		{name: "smoke 2", args: args{root: &TreeNode{Val: 1, Left: nil, Right: nil}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countNodes(test.args.root); got != test.want {
				t.Errorf("countNodes() = %v, want = %v", got, test.want)
			}
		})
	}
}
