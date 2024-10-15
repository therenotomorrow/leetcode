package golang

import (
	"reflect"
	"slices"
	"testing"
)

func TestRestoreArray(t *testing.T) {
	t.Parallel()

	type args struct {
		adjacentPairs [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{adjacentPairs: [][]int{{2, 1}, {3, 4}, {3, 2}}}, want: []int{1, 2, 3, 4}},
		{name: "smoke 2", args: args{adjacentPairs: [][]int{{4, -2}, {1, 4}, {-3, 1}}}, want: []int{-2, 4, 1, -3}},
		{name: "smoke 3", args: args{adjacentPairs: [][]int{{100000, -100000}}}, want: []int{100000, -100000}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := restoreArray(test.args.adjacentPairs)

			// because there could be different queue of traversal
			if got[0] != test.want[0] {
				slices.Reverse(test.want)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("restoreArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
