package golang

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
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
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{Val: 9, Left: nil, Right: nil},
				},
			}},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
					Right: &TreeNode{Val: 6, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 1, Left: nil, Right: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
				Right: &TreeNode{Val: 1, Left: nil, Right: nil},
			},
		},
		{
			name: "smoke 3",
			args: args{root: nil},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := invertTree(test.args.root); !reflect.DeepEqual(got, test.want) {
				t.Errorf("invertTree() = %v, want = %v", got, test.want)
			}
		})
	}
}
