package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetSneakyNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 1, 0}}, want: []int{0, 1}},
		{name: "smoke 2", args: args{nums: []int{0, 3, 2, 1, 3, 2}}, want: []int{2, 3}},
		{name: "smoke 3", args: args{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}}, want: []int{4, 5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := getSneakyNumbers(test.args.nums)

			sort.Ints(got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("getSneakyNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
