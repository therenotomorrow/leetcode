package maxLevelSum

import "testing"

func Test_maxLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: -8},
				},
				Right: &TreeNode{Val: 0},
			}},
			want: 2,
		},
		{
			args: args{root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			want: 2,
		},
		{
			args: args{root: &TreeNode{
				Val: -100,
				Left: &TreeNode{
					Val:   -200,
					Left:  &TreeNode{Val: -20},
					Right: &TreeNode{Val: -5},
				},
				Right: &TreeNode{
					Val:  -300,
					Left: &TreeNode{Val: -10},
				},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
