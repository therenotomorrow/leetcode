package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestMaximumAverageSubtree(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "smoke 1",
			args: args{root: &structs.TreeNode{
				Val:   5,
				Left:  &structs.TreeNode{Val: 6, Left: nil, Right: nil},
				Right: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
			}},
			want: 6.0,
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{
				Val:   0,
				Left:  nil,
				Right: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
			}},
			want: 1.0,
		},
		// 0,6,5 | 3,2,null,null | null,4,null,null | null,1
		{
			name: "test 15: wrong answer",
			args: args{root: &structs.TreeNode{
				Val: 0,
				Left: &structs.TreeNode{
					Val: 6,
					Left: &structs.TreeNode{
						Val:  3,
						Left: nil,
						Right: &structs.TreeNode{
							Val:   4,
							Left:  nil,
							Right: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
						},
					},
					Right: &structs.TreeNode{Val: 2, Left: nil, Right: nil},
				},
				Right: &structs.TreeNode{Val: 5, Left: nil, Right: nil},
			}},
			want: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumAverageSubtree(tt.args.root); got != tt.want {
				t.Errorf("maximumAverageSubtree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
