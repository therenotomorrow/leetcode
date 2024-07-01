package golang

import (
	"testing"
)

func TestMaxAncestorDiff(t *testing.T) {
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
					Val: 8,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   6,
							Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
							Right: &TreeNode{Val: 7, Left: nil, Right: nil},
						},
					},
					Right: &TreeNode{
						Val:  10,
						Left: nil,
						Right: &TreeNode{
							Val:   14,
							Left:  &TreeNode{Val: 13, Left: nil, Right: nil},
							Right: nil,
						},
					},
				},
			},
			want: 7,
		},
		{
			name: "smoke 2",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   0,
							Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
							Right: nil,
						},
					},
				},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxAncestorDiff(test.args.root); got != test.want {
				t.Errorf("maxAncestorDiff() = %v, want = %v", got, test.want)
			}
		})
	}
}
