package golang

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 1}, k: 3}, want: true},
		{name: "smoke 2", args: args{nums: []int{1, 0, 1, 1}, k: 1}, want: true},
		{name: "smoke 3", args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := containsNearbyDuplicate(test.args.nums, test.args.k); got != test.want {
				t.Errorf("containsNearbyDuplicate() = %v, want = %v", got, test.want)
			}
		})
	}
}
