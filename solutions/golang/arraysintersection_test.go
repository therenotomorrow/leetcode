package golang

import (
	"reflect"
	"testing"
)

func TestArraysIntersection(t *testing.T) {
	t.Parallel()

	type args struct {
		arr1 []int
		arr2 []int
		arr3 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{1, 2, 5, 7, 9},
				arr3: []int{1, 3, 4, 5, 8},
			},
			want: []int{1, 5},
		},
		{
			name: "smoke 2",
			args: args{
				arr1: []int{197, 418, 523, 876, 1356},
				arr2: []int{501, 880, 1593, 1710, 1870},
				arr3: []int{521, 682, 1337, 1395, 1764},
			},
			want: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := arraysIntersection(test.args.arr1, test.args.arr2, test.args.arr3); !reflect.DeepEqual(got, test.want) {
				t.Errorf("arraysIntersection() = %v, want = %v", got, test.want)
			}
		})
	}
}
