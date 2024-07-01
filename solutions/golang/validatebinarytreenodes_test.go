package golang

import "testing"

func TestValidateBinaryTreeNodes(t *testing.T) {
	t.Parallel()

	type args struct {
		n          int
		leftChild  []int
		rightChild []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 4, leftChild: []int{1, -1, 3, -1}, rightChild: []int{2, -1, -1, -1}}, want: true},
		{name: "smoke 2", args: args{n: 4, leftChild: []int{1, -1, 3, -1}, rightChild: []int{2, 3, -1, -1}}, want: false},
		{name: "smoke 3", args: args{n: 2, leftChild: []int{1, 0}, rightChild: []int{-1, -1}}, want: false},
		{
			name: "test 36: wrong answer",
			args: args{
				n:          6,
				leftChild:  []int{1, -1, -1, 4, -1, -1},
				rightChild: []int{2, -1, -1, 5, -1, -1},
			},
			want: false,
		},
		{
			name: "test 42: wrong answer",
			args: args{
				n:          3,
				leftChild:  []int{1, -1, -1},
				rightChild: []int{-1, -1, 1},
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := validateBinaryTreeNodes(test.args.n, test.args.leftChild, test.args.rightChild); got != test.want {
				t.Errorf("validateBinaryTreeNodes() = %v, want = %v", got, test.want)
			}
		})
	}
}
