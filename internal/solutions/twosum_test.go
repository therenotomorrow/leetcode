package solutions

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "smoke 2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "smoke 3", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			// You can return the answer in any order.
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want = %v", got, tt.want)
			}
		})
	}
}
