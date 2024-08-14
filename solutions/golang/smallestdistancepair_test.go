package golang

import "testing"

func TestSmallestDistancePair(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 1}, k: 1}, want: 0},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1}, k: 2}, want: 0},
		{name: "smoke 3", args: args{nums: []int{1, 6, 1}, k: 3}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := smallestDistancePair(test.args.nums, test.args.k); got != test.want {
				t.Errorf("smallestDistancePair() = %v, want = %v", got, test.want)
			}
		})
	}
}
