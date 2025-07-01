package golang

import "testing"

func TestLargestBSTSubtree(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 8, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			}},
			want: 3,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{Val: 2,
							Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: nil,
				},
			}},
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestBSTSubtree(test.args.root); got != test.want {
				t.Errorf("largestBSTSubtree() = %v, want = %v", got, test.want)
			}
		})
	}
}
