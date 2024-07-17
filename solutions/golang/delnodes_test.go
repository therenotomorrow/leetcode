package golang

import (
	"reflect"
	"slices"
	"testing"
)

func TestDelNodes(t *testing.T) {
	t.Parallel()

	type args struct {
		root     *TreeNode
		toDelete []int
	}

	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: Smoke1,
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
			}, toDelete: []int{3, 5}},
			want: nil,
		},
		{
			name: Smoke2,
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 4, Left: nil, Right: nil},
			}, toDelete: []int{3}},
			want: nil,
		},
	}

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			test.want = []*TreeNode{test.args.root, test.args.root.Right.Left, test.args.root.Right.Right}

		case Smoke2:
			test.want = []*TreeNode{test.args.root}
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := delNodes(test.args.root, test.args.toDelete)

			slices.SortStableFunc(got, func(a, b *TreeNode) int {
				return a.Val - b.Val
			})

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("delNodes() = %v, want = %v", got, test.want)
			}
		})
	}
}
