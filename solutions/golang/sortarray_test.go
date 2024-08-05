package golang

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 2, 3, 1}}, want: []int{1, 2, 3, 5}},
		{name: "smoke 2", args: args{nums: []int{5, 1, 1, 2, 0, 0}}, want: []int{0, 0, 1, 1, 2, 5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sortArray(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("sortArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
