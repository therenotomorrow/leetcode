package golang

import "testing"

func TestGetDirections(t *testing.T) {
	t.Parallel()

	type args struct {
		root       *TreeNode
		startValue int
		destValue  int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
			}, startValue: 3, destValue: 6},
			want: "UURL",
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: nil,
			}, startValue: 2, destValue: 1},
			want: "L",
		},
		{
			name: "test 237: wrong answer",
			args: args{root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: nil,
			}, startValue: 2, destValue: 1},
			want: "U",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getDirections(test.args.root, test.args.startValue, test.args.destValue); got != test.want {
				t.Errorf("getDirections() = %v, want = %v", got, test.want)
			}
		})
	}
}
