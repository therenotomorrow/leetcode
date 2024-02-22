package golang

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			}},
			want: []int{3, 2, 1},
		},
		{name: "smoke 2", args: args{root: nil}, want: []int{}},
		{name: "smoke 3", args: args{root: &TreeNode{Val: 1, Left: nil, Right: nil}}, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want = %v", got, tt.want)
			}
		})
	}
}
