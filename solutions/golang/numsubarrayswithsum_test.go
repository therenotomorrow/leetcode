package golang

import "testing"

func TestNumSubarraysWithSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		goal int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 0, 1, 0, 1}, goal: 2}, want: 4},
		{name: "smoke 2", args: args{nums: []int{0, 0, 0, 0, 0}, goal: 0}, want: 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numSubarraysWithSum(test.args.nums, test.args.goal); got != test.want {
				t.Errorf("numSubarraysWithSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
