package golang

import (
	"reflect"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
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
					Right: nil,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
			}},
			want: nil,
		},
		{
			name: Smoke2,
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 1, Left: nil, Right: nil},
			}},
			want: nil,
		},
		{
			name: Smoke3,
			args: args{root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			}},
			want: nil,
		},
	}

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			test.want = []*TreeNode{test.args.root.Right.Left.Left, test.args.root.Right.Left}

		case Smoke2:
			test.want = []*TreeNode{test.args.root.Right}

		case Smoke3:
			test.want = []*TreeNode{test.args.root.Right.Left, test.args.root.Right}
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findDuplicateSubtrees(test.args.root); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want = %v", got, test.want)
			}
		})
	}
}
