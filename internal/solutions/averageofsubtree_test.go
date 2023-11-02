package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"testing"
)

func TestAverageOfSubtree(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1", args: args{
				root: &structs.TreeNode{
					Val: 4,
					Left: &structs.TreeNode{
						Val:   8,
						Left:  &structs.TreeNode{Val: 0},
						Right: &structs.TreeNode{Val: 1},
					},
					Right: &structs.TreeNode{
						Val:   5,
						Right: &structs.TreeNode{Val: 6},
					},
				},
			},
			want: 5,
		},
		{name: "smoke 2", args: args{root: &structs.TreeNode{Val: 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfSubtree(tt.args.root); got != tt.want {
				t.Errorf("averageOfSubtree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
