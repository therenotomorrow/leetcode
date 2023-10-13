package longestSubarray

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{1, 1, 0, 1}}, want: 3},
		{args: args{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, want: 5},
		{args: args{nums: []int{1, 1, 1}}, want: 2},
		{args: args{nums: []int{1, 1, 0, 0, 1, 1, 1, 0, 1}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubarray2(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{8, 2, 4, 7}, limit: 4}, want: 2},
		{args: args{nums: []int{10, 1, 2, 4, 7, 2}, limit: 5}, want: 4},
		{args: args{nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray2(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
