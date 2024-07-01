package golang

import "testing"

func TestLongestOnes(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2}, want: 6},
		{name: "smoke 2", args: args{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3}, want: 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestOnes(test.args.nums, test.args.k); got != test.want {
				t.Errorf("longestOnes() = %v, want = %v", got, test.want)
			}
		})
	}
}
