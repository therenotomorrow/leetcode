package solutions

import (
	"reflect"
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestPreorderTraversal(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{root: &structs.TreeNode{
				Val:  1,
				Left: nil,
				Right: &structs.TreeNode{
					Val:   2,
					Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			}},
			want: []int{1, 2, 3},
		},
		{name: "smoke 2", args: args{root: nil}, want: []int{}},
		{name: "smoke 3", args: args{root: &structs.TreeNode{Val: 1, Left: nil, Right: nil}}, want: []int{1}},
		{
			name: "test 35: wrong answer",
			args: args{root: &structs.TreeNode{
				Val:   3,
				Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &structs.TreeNode{Val: 2, Left: nil, Right: nil},
			}},
			want: []int{3, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want = %v", got, tt.want)
			}
		})
	}
}
