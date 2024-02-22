package golang

import "testing"

func TestMinimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 5, 0, 3, 5}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{0}}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want = %v", got, tt.want)
			}
		})
	}
}
