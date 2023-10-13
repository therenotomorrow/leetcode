package closestValue

import "testing"

func Test_closestValue(t *testing.T) {
	type args struct {
		root   *TreeNode
		target float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 5},
				},
				target: 3.714286,
			},
			want: 4,
		},
		{
			args: args{root: &TreeNode{Val: 1}, target: 4.428571},
			want: 1,
		},
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				target: 3.5,
			},
			want: 3,
		},
		{
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				target: 3.428571,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestValue(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("closestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
