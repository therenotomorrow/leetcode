package solutions

import (
	"reflect"
	"slices"
	"testing"
)

func TestRestoreArray(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := restoreArray(tt.args.adjacentPairs)

			// because there could be different order of traversal
			if got[0] != tt.want[0] {
				slices.Reverse(tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreArray() = %v, want = %v", got, tt.want)
			}
		})
	}
}
