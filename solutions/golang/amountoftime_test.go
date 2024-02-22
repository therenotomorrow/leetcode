package golang

import (
	"testing"
)

func TestAmountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
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
					Val: 1,
					Left: &TreeNode{
						Val:  5,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
							Right: &TreeNode{Val: 2, Left: nil, Right: nil},
						},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 10, Left: nil, Right: nil},
						Right: &TreeNode{Val: 6, Left: nil, Right: nil},
					},
				},
				start: 3,
			},
			want: 4,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{Val: 1, Left: nil, Right: nil}, start: 1},
			want: 0,
		},
		{
			name: "test 66: wrong answer",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val:   4,
								Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
								Right: nil,
							},
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				start: 2,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amountOfTime(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTime() = %v, want = %v", got, tt.want)
			}
		})
	}
}
