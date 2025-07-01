package golang

import "testing"

func TestPartitionArray(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{3, 6, 1, 2, 5}, k: 2}, want: 2},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3}, k: 1}, want: 2},
		{name: "smoke 3", args: args{nums: []int{2, 2, 4, 5}, k: 0}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := partitionArray(test.args.nums, test.args.k); got != test.want {
				t.Errorf("partitionArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
