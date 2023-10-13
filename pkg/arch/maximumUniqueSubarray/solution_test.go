package maximumUniqueSubarray

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{4, 2, 6, 4, 5, 6}}, want: 17},
		{args: args{nums: []int{4, 2, 4, 5, 6}}, want: 17},
		{args: args{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}}, want: 8},
		{args: args{nums: []int{187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970, 306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434}}, want: 16911},
		//                                               ^                                                      ^                             ^                             ^                                                                      ^
		{args: args{nums: []int{10000, 1, 10000, 1, 1, 1, 1, 1, 1}}, want: 10001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.args.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
