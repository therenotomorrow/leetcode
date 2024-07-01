package golang

import "testing"

func TestMaxFrequency(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{1, 2, 4}, k: 5}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 4, 8, 13}, k: 5}, want: 2},
		{name: "smoke 3", args: args{nums: []int{3, 9, 6}, k: 2}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxFrequency(test.args.nums, test.args.k); got != test.want {
				t.Errorf("maxFrequency() = %v, want = %v", got, test.want)
			}
		})
	}
}
