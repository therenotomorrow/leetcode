package longestZigZag

import "testing"

func Test_longestZigZag(t *testing.T) {
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
				Right: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val:   1,
								Right: &TreeNode{Val: 1},
							},
						},
						Right: &TreeNode{Val: 1},
					},
				},
			}},
			want: 3,
		},
		{
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   1,
							Right: &TreeNode{Val: 1},
						},
						Right: &TreeNode{Val: 1},
					},
				},
				Right: &TreeNode{Val: 1},
			}},
			want: 4,
		},
		{
			args: args{root: &TreeNode{Val: 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestZigZag(tt.args.root); got != tt.want {
				t.Errorf("longestZigZag() = %v, want %v", got, tt.want)
			}
		})
	}
}
