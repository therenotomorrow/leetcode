package golang

import "testing"

type BinaryMatrixImpl struct {
	data [][]int
}

func (mat *BinaryMatrixImpl) Get(row int, col int) int {
	return mat.data[row][col]
}

func (mat *BinaryMatrixImpl) Dimensions() []int {
	return []int{len(mat.data), len(mat.data[0])}
}

func TestLeftMostColumnWithOne(t *testing.T) {
	t.Parallel()

	type args struct {
		binaryMatrix BinaryMatrix
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{binaryMatrix: &BinaryMatrixImpl{data: [][]int{{0, 0}, {1, 1}}}}, want: 0},
		{name: "smoke 2", args: args{binaryMatrix: &BinaryMatrixImpl{data: [][]int{{0, 0}, {0, 1}}}}, want: 1},
		{name: "smoke 3", args: args{binaryMatrix: &BinaryMatrixImpl{data: [][]int{{0, 0}, {0, 0}}}}, want: -1},
		{
			name: "test 6: wrong answer",
			args: args{binaryMatrix: &BinaryMatrixImpl{data: [][]int{
				{0, 0, 0, 1},
				{0, 0, 1, 1},
				{0, 1, 1, 1},
			}}},
			want: 1,
		},
		{
			name: "test 7: wrong answer",
			args: args{binaryMatrix: &BinaryMatrixImpl{data: [][]int{
				{1, 1, 1, 1, 1},
				{0, 0, 0, 1, 1},
				{0, 0, 1, 1, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
			}}},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := leftMostColumnWithOne(test.args.binaryMatrix); got != test.want {
				t.Errorf("leftMostColumnWithOne() = %v, want = %v", got, test.want)
			}
		})
	}
}
