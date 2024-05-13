package golang

import "testing"

func TestCountPairs(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("countPairs() = %v, want = %v", got, tt.want)
			}
		})
	}
}
