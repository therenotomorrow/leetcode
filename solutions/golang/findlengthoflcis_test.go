package golang

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 5, 4, 7}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{2, 2, 2, 2, 2}}, want: 1},
		{name: "test 13: wrong answer", args: args{nums: []int{1, 3, 5, 7}}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findLengthOfLCIS(test.args.nums); got != test.want {
				t.Errorf("findLengthOfLCIS() = %v, want = %v", got, test.want)
			}
		})
	}
}
