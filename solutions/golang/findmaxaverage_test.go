package golang

import "testing"

func TestFindMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "smoke 1", args: args{nums: []int{1, 12, -5, -6, 50, 3}, k: 4}, want: 12.75},
		{name: "smoke 2", args: args{nums: []int{5}, k: 1}, want: 5.0},
		{name: "test 21: wrong answer", args: args{nums: []int{0, 4, 0, 3, 2}, k: 1}, want: 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want = %v", got, tt.want)
			}
		})
	}
}
