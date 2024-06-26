package golang

import (
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
		low  int
		high int
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
					Val: 10,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 7, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: &TreeNode{Val: 18, Left: nil, Right: nil},
					},
				},
				low:  7,
				high: 15,
			},
			want: 32,
		},
		{
			name: "smoke 2",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   3,
							Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 13, Left: nil, Right: nil},
						Right: &TreeNode{Val: 18, Left: nil, Right: nil},
					},
				},
				low:  6,
				high: 10,
			},
			want: 23,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := rangeSumBST(test.args.root, test.args.low, test.args.high); got != test.want {
				t.Errorf("rangeSumBST() = %v, want = %v", got, test.want)
			}
		})
	}
}
