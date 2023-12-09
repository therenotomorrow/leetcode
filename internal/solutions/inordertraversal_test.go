package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
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
				Val: 1,
				Right: &structs.TreeNode{
					Val:  2,
					Left: &structs.TreeNode{Val: 3},
				},
			}},
			want: []int{1, 3, 2},
		},
		{name: "smoke 2", args: args{root: nil}, want: []int{}},
		{name: "smoke 3", args: args{root: &structs.TreeNode{Val: 1}}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want = %v", got, tt.want)
			}
		})
	}
}
