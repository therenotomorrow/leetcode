package golang

import (
	"reflect"
	"testing"
)

func TestRemoveLeafNodes(t *testing.T) {
	t.Parallel()

	type args struct {
		root   *TreeNode
		target int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "smoke 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: &TreeNode{Val: 4, Left: nil, Right: nil},
					},
				},
				target: 2,
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				target: 3,
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
				Right: nil,
			},
		},
		{
			name: "smoke 3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				target: 2,
			},
			want: &TreeNode{Val: 1, Left: nil, Right: nil},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := removeLeafNodes(test.args.root, test.args.target); !reflect.DeepEqual(got, test.want) {
				t.Errorf("removeLeafNodes() = %v, want = %v", got, test.want)
			}
		})
	}
}
