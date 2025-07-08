package golang

import (
	"reflect"
	"testing"
)

func TestInorderSuccessor(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
		p    *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: Smoke1,
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				p: nil,
			},
			want: nil,
		},
		{
			name: Smoke2,
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
					Right: &TreeNode{Val: 6, Left: nil, Right: nil},
				},
				p: nil,
			},
			want: nil,
		},
	}

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			test.args.p = test.args.root.Left
			test.want = test.args.root
		case Smoke2:
			test.args.p = test.args.root.Right
			test.want = nil
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := inorderSuccessor(test.args.root, test.args.p); !reflect.DeepEqual(got, test.want) {
				t.Errorf("inorderSuccessor() = %v, want = %v", got, test.want)
			}
		})
	}
}
