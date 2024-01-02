package solutions

import (
	"reflect"
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestFindMode(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1", args: args{
				root: &structs.TreeNode{
					Val:  1,
					Left: nil,
					Right: &structs.TreeNode{
						Val:   2,
						Left:  &structs.TreeNode{Val: 2, Left: nil, Right: nil},
						Right: nil,
					},
				},
			},
			want: []int{2},
		},
		{name: "smoke 2", args: args{root: &structs.TreeNode{Val: 0, Left: nil, Right: nil}}, want: []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want = %v", got, tt.want)
			}
		})
	}
}
