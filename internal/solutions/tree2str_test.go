package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestTree2str(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{root: &structs.TreeNode{
				Val: 1,
				Left: &structs.TreeNode{
					Val:   2,
					Left:  &structs.TreeNode{Val: 4, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &structs.TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: "1(2(4))(3)",
		},
		{
			name: "smoke 2",
			args: args{root: &structs.TreeNode{
				Val: 1,
				Left: &structs.TreeNode{
					Val:   2,
					Left:  nil,
					Right: &structs.TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &structs.TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: "1(2()(4))(3)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want = %v", got, tt.want)
			}
		})
	}
}
