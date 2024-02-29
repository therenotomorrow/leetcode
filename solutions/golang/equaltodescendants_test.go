package golang

import "testing"

func TestEqualToDescendants(t *testing.T) {
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
					Val:   3,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 1, Left: nil, Right: nil},
				},
				Right: &TreeNode{Val: 4, Left: nil, Right: nil},
			}},
			want: 2,
		},
		{
			name: "smoke 2",
			args: args{root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: nil,
				},
				Right: nil,
			}},
			want: 0,
		},
		{
			name: "smoke 3",
			args: args{root: &TreeNode{Val: 0, Left: nil, Right: nil}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalToDescendants(tt.args.root); got != tt.want {
				t.Errorf("equalToDescendants() = %v, want = %v", got, tt.want)
			}
		})
	}
}
