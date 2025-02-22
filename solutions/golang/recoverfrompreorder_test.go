package golang

import (
	"reflect"
	"testing"
)

func TestRecoverFromPreorder(t *testing.T) {
	t.Parallel()

	type args struct {
		traversal string
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "smoke 1",
			args: args{traversal: "1-2--3--4-5--6--7"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{traversal: "1-2--3---4-5--6---7"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			name: "smoke 3",
			args: args{traversal: "1-401--349---90--88"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 401,
					Left: &TreeNode{
						Val:   349,
						Left:  &TreeNode{Val: 90, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   88,
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := recoverFromPreorder(test.args.traversal); !reflect.DeepEqual(got, test.want) {
				t.Errorf("recoverFromPreorder() = %v, want = %v", got, test.want)
			}
		})
	}
}
