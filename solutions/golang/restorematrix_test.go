package golang

import (
	"reflect"
	"testing"
)

func TestRestoreMatrix(t *testing.T) {
	t.Parallel()

	type args struct {
		rowSum []int
		colSum []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{rowSum: []int{3, 8}, colSum: []int{4, 7}},
			want: [][]int{{3, 0}, {1, 7}},
		},
		{
			name: "smoke 2",
			args: args{rowSum: []int{5, 7, 10}, colSum: []int{8, 6, 8}},
			want: [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := restoreMatrix(test.args.rowSum, test.args.colSum); !reflect.DeepEqual(got, test.want) {
				t.Errorf("restoreMatrix() = %v, want = %v", got, test.want)
			}
		})
	}
}
