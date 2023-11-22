package solutions

import "testing"

func TestMaxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 4}, k: 5}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 4, 8, 13}, k: 5}, want: 2},
		{name: "smoke 3", args: args{nums: []int{3, 9, 6}, k: 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want = %v", got, tt.want)
			}
		})
	}
}
