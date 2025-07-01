package golang

import (
	"reflect"
	"testing"
)

func TestTreeToDoublyList(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: Smoke1,
			args: args{root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 5, Left: nil, Right: nil},
			}},
			want: &TreeNode{
				Val: 1, Left: nil, Right: &TreeNode{
					Val: 2, Left: nil, Right: &TreeNode{
						Val: 3, Left: nil, Right: &TreeNode{
							Val: 4, Left: nil, Right: &TreeNode{
								Val: 5, Left: nil, Right: nil,
							},
						},
					},
				},
			},
		},
		{
			name: Smoke2,
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: &TreeNode{
				Val: 1, Left: nil, Right: &TreeNode{
					Val: 2, Left: nil, Right: &TreeNode{
						Val: 3, Left: nil, Right: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			test.want.Left = test.want.Right.Right.Right.Right

			test.want.Right.Left = test.want
			test.want.Right.Right.Left = test.want.Right
			test.want.Right.Right.Right.Left = test.want.Right.Right
			test.want.Right.Right.Right.Right.Left = test.want.Right.Right.Right

			test.want.Right.Right.Right.Right.Right = test.want

		case Smoke2:
			test.want.Left = test.want.Right.Right

			test.want.Right.Left = test.want
			test.want.Right.Right.Left = test.want.Right

			test.want.Right.Right.Right = test.want
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := treeToDoublyList(test.args.root); !reflect.DeepEqual(got, test.want) {
				t.Errorf("treeToDoublyList() = %v, want = %v", got, test.want)
			}
		})
	}
}
