package solutions

import (
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want = %v", got, tt.want)
			}
		})
	}
}
