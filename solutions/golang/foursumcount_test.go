package golang

import "testing"

func TestFourSumCount(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{-2, -1},
				nums3: []int{-1, 2},
				nums4: []int{0, 2},
			},
			want: 2,
		},
		{
			name: "smoke 2",
			args: args{
				nums1: []int{0},
				nums2: []int{0},
				nums3: []int{0},
				nums4: []int{0},
			},
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := fourSumCount(test.args.nums1, test.args.nums2, test.args.nums3, test.args.nums4); got != test.want {
				t.Errorf("fourSumCount() = %v, want = %v", got, test.want)
			}
		})
	}
}
