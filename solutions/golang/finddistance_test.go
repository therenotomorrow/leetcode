package golang

import "testing"

func TestFindDistance(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
		p    int
		q    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
					Right: &TreeNode{Val: 8, Left: nil, Right: nil},
				},
			}, p: 5, q: 0},
			want: 3,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
					Right: &TreeNode{Val: 8, Left: nil, Right: nil},
				},
			}, p: 5, q: 7},
			want: 2,
		},
		{
			name: "smoke 3",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
					Right: &TreeNode{Val: 8, Left: nil, Right: nil},
				},
			}, p: 5, q: 5},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findDistance(test.args.root, test.args.p, test.args.q); got != test.want {
				t.Errorf("findDistance() = %v, want = %v", got, test.want)
			}
		})
	}
}
