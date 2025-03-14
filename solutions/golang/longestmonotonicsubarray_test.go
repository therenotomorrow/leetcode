package golang

import "testing"

func TestLongestMonotonicSubarray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 4, 3, 3, 2}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{3, 3, 3, 3}}, want: 1},
		{name: "smoke 3", args: args{nums: []int{3, 2, 1}}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestMonotonicSubarray(test.args.nums); got != test.want {
				t.Errorf("longestMonotonicSubarray() = %v, want = %v", got, test.want)
			}
		})
	}
}
