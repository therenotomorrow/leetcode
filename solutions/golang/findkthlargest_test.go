package golang

import "testing"

func TestFindKthLargest(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, want: 5},
		{name: "smoke 2", args: args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findKthLargest(test.args.nums, test.args.k); got != test.want {
				t.Errorf("findKthLargest() = %v, want = %v", got, test.want)
			}
		})
	}
}
