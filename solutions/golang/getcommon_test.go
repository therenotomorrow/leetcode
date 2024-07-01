package golang

import "testing"

func TestGetCommon(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums1: []int{1, 2, 3}, nums2: []int{2, 4}}, want: 2},
		{name: "smoke 2", args: args{nums1: []int{1, 2, 3, 6}, nums2: []int{2, 3, 4, 5}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getCommon(test.args.nums1, test.args.nums2); got != test.want {
				t.Errorf("getCommon() = %v, want = %v", got, test.want)
			}
		})
	}
}
