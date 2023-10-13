package twoSum

import (
	"reflect"
	"sort"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
