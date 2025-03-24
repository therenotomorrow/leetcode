package golang

import "testing"

func TestLongestNiceSubarray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 8, 48, 10}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{3, 1, 5, 11, 13}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestNiceSubarray(test.args.nums); got != test.want {
				t.Errorf("longestNiceSubarray() = %v, want = %v", got, test.want)
			}
		})
	}
}
