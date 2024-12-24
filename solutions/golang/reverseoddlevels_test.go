package golang

import (
	"reflect"
	"testing"
)

func TestReverseOddLevels(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 8, Left: nil, Right: nil},
					Right: &TreeNode{Val: 13, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 21, Left: nil, Right: nil},
					Right: &TreeNode{Val: 34, Left: nil, Right: nil},
				},
			}},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 8, Left: nil, Right: nil},
					Right: &TreeNode{Val: 13, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 21, Left: nil, Right: nil},
					Right: &TreeNode{Val: 34, Left: nil, Right: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 13, Left: nil, Right: nil},
				Right: &TreeNode{Val: 11, Left: nil, Right: nil},
			}},
			want: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 11, Left: nil, Right: nil},
				Right: &TreeNode{Val: 13, Left: nil, Right: nil},
			},
		},
		{
			name: "smoke 3",
			args: args{root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
				},
			}},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: &TreeNode{Val: 2, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &TreeNode{Val: 1, Left: nil, Right: nil},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := reverseOddLevels(test.args.root)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("reverseOddLevels() = %v, want = %v", got, test.want)
			}
		})
	}
}
