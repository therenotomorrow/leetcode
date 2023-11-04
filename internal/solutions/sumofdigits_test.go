package solutions

import "testing"

func TestSumOfDigits(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{34, 23, 1, 24, 75, 33, 54, 8}}, want: 0},
		{name: "smoke 2", args: args{nums: []int{99, 77, 33, 66, 55}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDigits(tt.args.nums); got != tt.want {
				t.Errorf("sumOfDigits() = %v, want = %v", got, tt.want)
			}
		})
	}
}
