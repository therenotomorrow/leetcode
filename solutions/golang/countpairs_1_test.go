package golang

import "testing"

func TestCountPairs1(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{nums1: []int{2, 1, 2, 1}, nums2: []int{1, 2, 1, 2}}, want: 1},
		{name: "smoke 2", args: args{nums1: []int{1, 10, 6, 2}, nums2: []int{1, 4, 1, 5}}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countPairs1(test.args.nums1, test.args.nums2); got != test.want {
				t.Errorf("countPairs1() = %v, want = %v", got, test.want)
			}
		})
	}
}
