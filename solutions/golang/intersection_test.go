package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 3}}, want: []int{2}},
		{name: "smoke 2", args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := intersection(test.args.nums1, test.args.nums2)

			sort.Ints(got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("intersection() = %v, want = %v", got, test.want)
			}
		})
	}
}
