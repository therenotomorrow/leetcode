package golang

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "smoke 2", args: args{nums: []int{1}}, want: 1},
		{name: "smoke 3", args: args{nums: []int{5, 4, -1, 7, 8}}, want: 23},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxSubArray(test.args.nums); got != test.want {
				t.Errorf("maxSubArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
