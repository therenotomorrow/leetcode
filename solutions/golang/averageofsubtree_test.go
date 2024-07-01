package golang

import (
	"testing"
)

func TestAverageOfSubtree(t *testing.T) {
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
			name: "smoke 1", args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: &TreeNode{Val: 6, Left: nil, Right: nil},
					},
				},
			},
			want: 5,
		},
		{name: "smoke 2", args: args{root: &TreeNode{Val: 1, Left: nil, Right: nil}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := averageOfSubtree(test.args.root); got != test.want {
				t.Errorf("averageOfSubtree() = %v, want = %v", got, test.want)
			}
		})
	}
}
