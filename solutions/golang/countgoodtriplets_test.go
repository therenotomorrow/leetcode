package golang

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{3, 0, 1, 1, 9, 7}, a: 7, b: 2, c: 3}, want: 4},
		{name: "smoke 2", args: args{arr: []int{1, 1, 2, 2, 3}, a: 0, b: 0, c: 1}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countGoodTriplets(test.args.arr, test.args.a, test.args.b, test.args.c); got != test.want {
				t.Errorf("countGoodTriplets() = %v, want = %v", got, test.want)
			}
		})
	}
}
