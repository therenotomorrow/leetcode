package golang

import "testing"

func TestLengthOfLIS(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{0, 1, 0, 3, 2, 3}}, want: 4},
		{name: "smoke 3", args: args{nums: []int{7, 7, 7, 7, 7, 7, 7}}, want: 1},
		{name: "test 35: wrong answer", args: args{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lengthOfLIS(test.args.nums); got != test.want {
				t.Errorf("lengthOfLIS() = %v, want = %v", got, test.want)
			}
		})
	}
}
