package solutions

import (
	"reflect"
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestBinaryTreePaths(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{root: &structs.TreeNode{
				Val: 1,
				Left: &structs.TreeNode{
					Val:   2,
					Left:  nil,
					Right: &structs.TreeNode{Val: 5, Left: nil, Right: nil},
				},
				Right: &structs.TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{Val: 1, Left: nil, Right: nil}},
			want: []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want = %v", got, tt.want)
			}
		})
	}
}
