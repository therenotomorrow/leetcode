package solutions

import "testing"

func TestArraySign(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-1, -2, -3, -4, 3, 2, 1}}, want: 1},
		{name: "smoke 2", args: args{nums: []int{1, 5, 0, 2, -3}}, want: 0},
		{name: "smoke 3", args: args{nums: []int{-1, 1, -1, 1, -1}}, want: -1},
		{name: "test 31: wrong answer", args: args{nums: []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraySign(tt.args.nums); got != tt.want {
				t.Errorf("arraySign() = %v, want = %v", got, tt.want)
			}
		})
	}
}
