package solutions

import (
	"reflect"
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestInvertTree(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want *structs.TreeNode
	}{
		{
			name: "smoke 1",
			args: args{root: &structs.TreeNode{
				Val: 4,
				Left: &structs.TreeNode{
					Val:   2,
					Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &structs.TreeNode{
					Val:   7,
					Left:  &structs.TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 9, Left: nil, Right: nil},
				},
			}},
			want: &structs.TreeNode{
				Val: 4,
				Left: &structs.TreeNode{
					Val:   7,
					Left:  &structs.TreeNode{Val: 9, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 6, Left: nil, Right: nil},
				},
				Right: &structs.TreeNode{
					Val:   2,
					Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{
				Val:   2,
				Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &structs.TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: &structs.TreeNode{
				Val:   2,
				Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
				Right: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
			},
		},
		{
			name: "smoke 3",
			args: args{root: nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
