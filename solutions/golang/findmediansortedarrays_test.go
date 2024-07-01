package golang

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "smoke 1", args: args{nums1: []int{1, 3}, nums2: []int{2}}, want: 2.0},
		{name: "smoke 2", args: args{nums1: []int{1, 2}, nums2: []int{3, 4}}, want: 2.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMedianSortedArrays(test.args.nums1, test.args.nums2); got != test.want {
				t.Errorf("findMedianSortedArrays() = %v, want = %v", got, test.want)
			}
		})
	}
}
