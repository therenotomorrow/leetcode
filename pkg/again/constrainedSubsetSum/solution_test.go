package constrainedSubsetSum

import "testing"

func Test_constrainedSubsetSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{10, 2, -10, 5, 20}, k: 2}, want: 37},
		{args: args{nums: []int{-1, -2, -3}, k: 1}, want: -1},
		{args: args{nums: []int{10, -2, -10, -5, 20}, k: 2}, want: 23},
		{args: args{nums: []int{-8269, 3217, -4023, -4138, -683, 6455, -3621, 9242, 4015, -3790}, k: 1}, want: 16091},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constrainedSubsetSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("constrainedSubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
