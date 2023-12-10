package solutions

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
		{name: "smoke 2", args: args{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
		{name: "smoke 3", args: args{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want = %v", got, tt.want)
			}
		})
	}
}
