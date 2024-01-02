package lowestCommonAncestor

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorCase1(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 6},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 4},
						},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  &TreeNode{Val: 0},
						Right: &TreeNode{Val: 8},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.args.p = tt.args.root.Left
		tt.args.q = tt.args.root.Right
		tt.want = tt.args.root

		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestorCase1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestorCase2(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 6},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 4},
						},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  &TreeNode{Val: 0},
						Right: &TreeNode{Val: 8},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.args.p = tt.args.root.Left
		tt.args.q = tt.args.root.Left.Right.Right
		tt.want = tt.args.root.Left

		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestorCase2() = %v, want %v", got, tt.want)
			}
		})
	}
}
