package golang

import (
	"reflect"
	"testing"
)

func TestOnesMinusZeros(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{grid: [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}},
			want: [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}},
		},
		{
			name: "smoke 2",
			args: args{grid: [][]int{{1, 1, 1}, {1, 1, 1}}},
			want: [][]int{{5, 5, 5}, {5, 5, 5}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := onesMinusZeros(test.args.grid); !reflect.DeepEqual(got, test.want) {
				t.Errorf("onesMinusZeros() = %v, want = %v", got, test.want)
			}
		})
	}
}
