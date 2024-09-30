package golang

import (
	"reflect"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	t.Parallel()

	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "smoke 1",
			args: args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
				root2: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: &TreeNode{Val: 7, Left: nil, Right: nil},
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{
				root1: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				root2: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: nil,
				},
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := mergeTrees(test.args.root1, test.args.root2); !reflect.DeepEqual(got, test.want) {
				t.Errorf("mergeTrees() = %v, want = %v", got, test.want)
			}
		})
	}
}
