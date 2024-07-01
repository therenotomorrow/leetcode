package golang

import (
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3}}, want: []int{1, 2}},
		{name: "smoke 2", args: args{nums: []int{1, 2, 4, 8}}, want: []int{1, 2, 4, 8}},
		{name: "test 21: wrong answer", args: args{nums: []int{4, 8, 10, 240}}, want: []int{4, 8, 240}},
		{name: "test 22: wrong answer", args: args{nums: []int{2, 3, 4, 9, 8}}, want: []int{2, 4, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestDivisibleSubset(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("largestDivisibleSubset() = %v, want = %v", got, test.want)
			}
		})
	}
}
