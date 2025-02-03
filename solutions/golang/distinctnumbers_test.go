package golang

import (
	"reflect"
	"testing"
)

func TestDistinctNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 2, 2, 1, 3}, k: 3}, want: []int{3, 2, 2, 2, 3}},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1, 1, 2, 3, 4}, k: 4}, want: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := distinctNumbers(test.args.nums, test.args.k); !reflect.DeepEqual(got, test.want) {
				t.Errorf("distinctNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
