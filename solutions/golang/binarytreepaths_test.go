package golang

import (
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{Val: 1, Left: nil, Right: nil}},
			want: []string{"1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := binaryTreePaths(test.args.root); !reflect.DeepEqual(got, test.want) {
				t.Errorf("binaryTreePaths() = %v, want = %v", got, test.want)
			}
		})
	}
}
