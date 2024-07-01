package golang

import (
	"reflect"
	"testing"
)

func TestFindMode(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1", args: args{
				root: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
						Right: nil,
					},
				},
			},
			want: []int{2},
		},
		{name: "smoke 2", args: args{root: &TreeNode{Val: 0, Left: nil, Right: nil}}, want: []int{0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMode(test.args.root); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findMode() = %v, want = %v", got, test.want)
			}
		})
	}
}
