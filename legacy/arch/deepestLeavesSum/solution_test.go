package deepestLeavesSum

import "testing"

func Test_deepestLeavesSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 7},
						},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:   6,
							Right: &TreeNode{Val: 8},
						},
					},
				},
			},
			want: 15,
		},
		{
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 9},
						},
						Right: &TreeNode{
							Val:   7,
							Left:  &TreeNode{Val: 1},
							Right: &TreeNode{Val: 4},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val:   3,
							Right: &TreeNode{Val: 5},
						},
					},
				},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
