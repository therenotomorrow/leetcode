package golang

import "testing"

func TestEvaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
					Right: &TreeNode{Val: 1, Left: nil, Right: nil},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{Val: 0, Left: nil, Right: nil}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateTree(tt.args.root); got != tt.want {
				t.Errorf("evaluateTree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
