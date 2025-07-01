package golang

import "testing"

func TestCheckEqualTree(t *testing.T) {
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
				Val:  5,
				Left: &TreeNode{Val: 10, Left: nil, Right: nil},
				Right: &TreeNode{
					Val:   10,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: &TreeNode{
					Val:   10,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 20, Left: nil, Right: nil},
				},
			}},
			want: false,
		},
		{
			name: "test 193: wrong answer",
			args: args{root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: -1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 1, Left: nil, Right: nil},
			}},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := checkEqualTree(test.args.root); got != test.want {
				t.Errorf("checkEqualTree() = %v, want = %v", got, test.want)
			}
		})
	}
}
