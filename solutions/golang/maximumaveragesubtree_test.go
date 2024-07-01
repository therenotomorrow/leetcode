package golang

import (
	"testing"
)

func TestMaximumAverageSubtree(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
				Right: &TreeNode{Val: 1, Left: nil, Right: nil},
			}},
			want: 6.0,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: &TreeNode{Val: 1, Left: nil, Right: nil},
			}},
			want: 1.0,
		},
		// 0,6,5 | 3,2,null,null | null,4,null,null | null,1
		{
			name: "test 15: wrong answer",
			args: args{root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: &TreeNode{Val: 1, Left: nil, Right: nil},
						},
					},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 5, Left: nil, Right: nil},
			}},
			want: 5.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumAverageSubtree(test.args.root); got != test.want {
				t.Errorf("maximumAverageSubtree() = %v, want = %v", got, test.want)
			}
		})
	}
}
