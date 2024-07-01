package golang

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	t.Parallel()

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "smoke 2",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := transpose(test.args.matrix); !reflect.DeepEqual(got, test.want) {
				t.Errorf("transpose() = %v, want = %v", got, test.want)
			}
		})
	}
}
