package golang

import "testing"

func TestFindMaxAverage(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMaxAverage(test.args.nums, test.args.k); got != test.want {
				t.Errorf("findMaxAverage() = %v, want = %v", got, test.want)
			}
		})
	}
}
