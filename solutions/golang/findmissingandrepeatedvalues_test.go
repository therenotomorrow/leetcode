package golang

import (
	"reflect"
	"testing"
)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{grid: [][]int{{1, 3}, {2, 2}}}, want: []int{2, 4}},
		{name: "smoke 2", args: args{grid: [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}}, want: []int{9, 5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMissingAndRepeatedValues(test.args.grid); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findMissingAndRepeatedValues() = %v, want = %v", got, test.want)
			}
		})
	}
}
