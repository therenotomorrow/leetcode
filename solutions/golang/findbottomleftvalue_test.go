package golang

import "testing"

func TestFindBottomLeftValue(t *testing.T) {
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
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			want: 1,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &TreeNode{Val: 6, Left: nil, Right: nil},
				},
			}},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want = %v", got, tt.want)
			}
		})
	}
}
