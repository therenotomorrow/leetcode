package deleteNode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 7},
					},
				},
				key: 3,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 7},
					},
				},
				key: 0,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			args: args{
				root: nil,
				key:  0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				fmt.Println(tt.want)

				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
