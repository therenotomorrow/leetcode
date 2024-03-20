package golang

import (
	"reflect"
	"testing"
)

func TestTwoOutOfThree(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums1: []int{1, 1, 3, 2}, nums2: []int{2, 3}, nums3: []int{3}}, want: []int{2, 3}},
		{name: "smoke 2", args: args{nums1: []int{3, 1}, nums2: []int{2, 3}, nums3: []int{1, 2}}, want: []int{1, 2, 3}},
		{name: "smoke 3", args: args{nums1: []int{1, 2, 2}, nums2: []int{4, 3, 3}, nums3: []int{5}}, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoOutOfThree(tt.args.nums1, tt.args.nums2, tt.args.nums3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOutOfThree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
