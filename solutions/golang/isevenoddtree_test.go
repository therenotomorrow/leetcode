package golang

import "testing"

func TestIsEvenOddTree(t *testing.T) {
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
				Val: 1,
				Left: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 12, Left: nil, Right: nil},
						Right: &TreeNode{Val: 8, Left: nil, Right: nil},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
					Right: nil,
				},
			}},
			want: false,
		},
		{
			name: "smoke 3",
			args: args{root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   9,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
					Right: nil,
				},
			}},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isEvenOddTree(test.args.root); got != test.want {
				t.Errorf("isEvenOddTree() = %v, want = %v", got, test.want)
			}
		})
	}
}
