package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestRangeSumBST(t *testing.T) {
	type args struct {
		root *structs.TreeNode
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
				root: &structs.TreeNode{
					Val: 10,
					Left: &structs.TreeNode{
						Val:   5,
						Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 7, Left: nil, Right: nil},
					},
					Right: &structs.TreeNode{
						Val:   15,
						Left:  nil,
						Right: &structs.TreeNode{Val: 18, Left: nil, Right: nil},
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
				root: &structs.TreeNode{
					Val: 10,
					Left: &structs.TreeNode{
						Val: 5,
						Left: &structs.TreeNode{
							Val:   3,
							Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
						Right: &structs.TreeNode{
							Val:   7,
							Left:  &structs.TreeNode{Val: 6, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &structs.TreeNode{
						Val:   15,
						Left:  &structs.TreeNode{Val: 13, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 18, Left: nil, Right: nil},
					},
				},
				low:  6,
				high: 10,
			},
			want: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want = %v", got, tt.want)
			}
		})
	}
}
