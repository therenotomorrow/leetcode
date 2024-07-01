package golang

import "testing"

func TestIsSameTree(t *testing.T) {
	t.Parallel()

	type args struct {
		p *TreeNode
		q *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				q: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: nil,
				},
				q: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: false,
		},
		{
			name: "smoke 3",
			args: args{
				p: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 1, Left: nil, Right: nil},
				},
				q: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isSameTree(test.args.p, test.args.q); got != test.want {
				t.Errorf("isSameTree() = %v, want = %v", got, test.want)
			}
		})
	}
}
