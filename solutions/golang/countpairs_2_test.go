package golang

import "testing"

func TestCountPairs2(t *testing.T) {
	t.Parallel()

	type args struct {
		root     *TreeNode
		distance int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			}, distance: 3},
			want: 1,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			}, distance: 3},
			want: 2,
		},
		{
			name: "smoke 3",
			args: args{root: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
				},
			}, distance: 3},
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countPairs2(test.args.root, test.args.distance); got != test.want {
				t.Errorf("countPairs2() = %v, want = %v", got, test.want)
			}
		})
	}
}
