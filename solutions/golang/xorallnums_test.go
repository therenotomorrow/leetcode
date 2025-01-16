package golang

import "testing"

func TestXorAllNums(t *testing.T) {
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
		{name: "smoke 1", args: args{nums1: []int{2, 1, 3}, nums2: []int{10, 2, 5, 0}}, want: 13},
		{name: "smoke 2", args: args{nums1: []int{1, 2}, nums2: []int{3, 4}}, want: 0},
		{
			name: "test 23: wrong answer",
			args: args{nums1: []int{8, 6, 29, 2, 26, 16, 15, 29}, nums2: []int{24, 12, 12}},
			want: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := xorAllNums(test.args.nums1, test.args.nums2); got != test.want {
				t.Errorf("xorAllNums() = %v, want = %v", got, test.want)
			}
		})
	}
}
