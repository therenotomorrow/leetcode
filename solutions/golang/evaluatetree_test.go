package golang

import "testing"

func TestEvaluateTree(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := evaluateTree(test.args.root); got != test.want {
				t.Errorf("evaluateTree() = %v, want = %v", got, test.want)
			}
		})
	}
}
