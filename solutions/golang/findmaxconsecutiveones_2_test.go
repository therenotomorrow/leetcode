package golang

import "testing"

func TestFindMaxConsecutiveOnes2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 0, 1, 1, 0}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{1, 0, 1, 1, 0, 1}}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes2(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
