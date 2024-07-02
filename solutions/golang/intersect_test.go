package golang

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
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
		{name: "smoke 1", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{name: "smoke 2", args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{9, 4}},
		{name: "test 33: wrong answer", args: args{nums1: []int{3, 1, 2}, nums2: []int{1, 1}}, want: []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := intersect(test.args.nums1, test.args.nums2); !reflect.DeepEqual(got, test.want) {
				t.Errorf("intersect() = %v, want = %v", got, test.want)
			}
		})
	}
}
