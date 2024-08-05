package golang

import (
	"reflect"
	"testing"
)

func TestSortJumbled(t *testing.T) {
	t.Parallel()

	type args struct {
		mapping []int
		nums    []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, nums: []int{991, 338, 38}},
			want: []int{338, 38, 991},
		},
		{
			name: "smoke 2",
			args: args{mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nums: []int{789, 456, 123}},
			want: []int{123, 456, 789},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sortJumbled(test.args.mapping, test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("sortJumbled() = %v, want = %v", got, test.want)
			}
		})
	}
}
