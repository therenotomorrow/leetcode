package golang

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}}, want: 2},
		{name: "smoke 2", args: args{target: 4, nums: []int{1, 4, 4}}, want: 1},
		{name: "smoke 3", args: args{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}}, want: 0},
		{name: "test 11: wrong answer", args: args{target: 11, nums: []int{1, 2, 3, 4, 5}}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want = %v", got, tt.want)
			}
		})
	}
}
