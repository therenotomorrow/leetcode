package golang

import (
	"reflect"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	t.Parallel()

	type args struct {
		descriptions [][]int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "smoke 1",
			args: args{descriptions: [][]int{
				{20, 15, 1},
				{20, 17, 0},
				{50, 20, 1},
				{50, 80, 0},
				{80, 19, 1},
			}},
			want: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &TreeNode{Val: 17, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   80,
					Left:  &TreeNode{Val: 19, Left: nil, Right: nil},
					Right: nil,
				},
			},
		},
		{
			name: "smoke 2",
			args: args{descriptions: [][]int{
				{1, 2, 1},
				{2, 3, 0},
				{3, 4, 1},
			}},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: nil,
					},
				},
				Right: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := createBinaryTree(test.args.descriptions); !reflect.DeepEqual(got, test.want) {
				t.Errorf("createBinaryTree() = %v, want = %v", got, test.want)
			}
		})
	}
}
