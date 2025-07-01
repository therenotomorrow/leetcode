package golang

import "testing"

func TestMaximumDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{7, 1, 5, 4}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{9, 4, 3, 2}}, want: -1},
		{name: "smoke 3", args: args{nums: []int{1, 5, 2, 10}}, want: 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumDifference(test.args.nums); got != test.want {
				t.Errorf("maximumDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
