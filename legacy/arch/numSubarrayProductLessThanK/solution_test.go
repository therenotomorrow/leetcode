package numSubarrayProductLessThanK

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{10, 5, 2, 6}, k: 100}, want: 8},
		{args: args{nums: []int{1, 2, 3}, k: 0}, want: 0},
		{args: args{nums: []int{1, 2, 3}, k: 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
