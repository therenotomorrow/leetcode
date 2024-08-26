package golang

import (
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	t.Parallel()

	type args struct {
		root *Node
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{Val: 5, Children: nil},
							{Val: 6, Children: nil},
						},
					},
					{Val: 2, Children: nil},
					{Val: 4, Children: nil},
				},
			}},
			want: []int{5, 6, 3, 2, 4, 1},
		},
		{
			name: "smoke 2",
			args: args{root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val:      2,
						Children: nil,
					},
					{
						Val: 3,
						Children: []*Node{
							{
								Val:      6,
								Children: nil,
							},
							{
								Val: 7,
								Children: []*Node{
									{
										Val: 11,
										Children: []*Node{
											{
												Val:      14,
												Children: nil,
											},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*Node{
							{
								Val: 8,
								Children: []*Node{
									{
										Val:      12,
										Children: nil,
									},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*Node{
							{
								Val: 9,
								Children: []*Node{
									{
										Val:      13,
										Children: nil,
									},
								},
							},
							{
								Val:      10,
								Children: nil,
							},
						},
					},
				},
			}},
			want: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := postorder(test.args.root); !reflect.DeepEqual(got, test.want) {
				t.Errorf("postorder() = %v, want = %v", got, test.want)
			}
		})
	}
}
