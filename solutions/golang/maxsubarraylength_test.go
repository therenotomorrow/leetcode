package golang

import "testing"

func TestMaxSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 1, 2, 3, 1, 2}, k: 2}, want: 6},
		{name: "smoke 2", args: args{nums: []int{1, 2, 1, 2, 1, 2, 1, 2}, k: 1}, want: 2},
		{name: "smoke 3", args: args{nums: []int{5, 5, 5, 5, 5, 5, 5}, k: 4}, want: 4},
		{name: "test 730: wrong answer", args: args{nums: []int{3, 1, 1}, k: 1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarrayLength() = %v, want = %v", got, tt.want)
			}
		})
	}
}
