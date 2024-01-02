package insertIntoBST

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 7},
				},
				val: 5,
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		{
			// [40,20,60,  10,30,50,70, ]
			args: args{
				root: &TreeNode{
					Val: 40,
					Left: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 10},
						Right: &TreeNode{Val: 30},
					},
					Right: &TreeNode{
						Val:   60,
						Left:  &TreeNode{Val: 50},
						Right: &TreeNode{Val: 70},
					},
				},
				val: 25,
			},
			want: &TreeNode{
				Val: 40,
				Left: &TreeNode{
					Val:  20,
					Left: &TreeNode{Val: 10},
					Right: &TreeNode{
						Val:  30,
						Left: &TreeNode{Val: 25},
					},
				},
				Right: &TreeNode{
					Val:   60,
					Left:  &TreeNode{Val: 50},
					Right: &TreeNode{Val: 70},
				},
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				val: 5,
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
