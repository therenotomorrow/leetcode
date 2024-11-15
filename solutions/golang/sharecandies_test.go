package golang

import (
	"testing"
)

func TestShareCandies(t *testing.T) {
	t.Parallel()

	type args struct {
		candies []int
		k       int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{candies: []int{1, 2, 2, 3, 4, 3}, k: 3}, want: 3},
		{name: "smoke 2", args: args{candies: []int{2, 2, 2, 2, 3, 3}, k: 2}, want: 2},
		{name: "smoke 3", args: args{candies: []int{2, 4, 5}, k: 0}, want: 3},
		{name: "test 32: wrong answer", args: args{candies: []int{2, 4, 5}, k: 3}, want: 0},
		{name: "test 49: wrong answer", args: args{candies: []int{3, 3, 3, 4}, k: 3}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := shareCandies(test.args.candies, test.args.k); got != test.want {
				t.Errorf("shareCandies() = %v, want = %v", got, test.want)
			}
		})
	}
}
