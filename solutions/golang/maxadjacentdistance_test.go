package golang

import "testing"

func TestMaxAdjacentDistance(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 4}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{-5, -10, -5}}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxAdjacentDistance(test.args.nums); got != test.want {
				t.Errorf("maxAdjacentDistance() = %v, want = %v", got, test.want)
			}
		})
	}
}
