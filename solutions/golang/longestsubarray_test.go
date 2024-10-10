package golang

import "testing"

func TestLongestSubarray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 3, 2, 2}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestSubarray(test.args.nums); got != test.want {
				t.Errorf("longestSubarray() = %v, want = %v", got, test.want)
			}
		})
	}
}