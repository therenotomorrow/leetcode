package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"testing"
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
				Left:  &structs.TreeNode{Val: 6},
				Right: &structs.TreeNode{Val: 1},
			}},
			want: 6.0,
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{
				Val:   0,
				Right: &structs.TreeNode{Val: 1},
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
						Val: 3,
						Right: &structs.TreeNode{
							Val:   4,
							Right: &structs.TreeNode{Val: 1},
						},
					},
					Right: &structs.TreeNode{Val: 2},
				},
				Right: &structs.TreeNode{Val: 5},
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
