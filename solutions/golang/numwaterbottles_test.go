package golang

import "testing"

func TestNumWaterBottles(t *testing.T) {
	t.Parallel()

	type args struct {
		numBottles  int
		numExchange int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{numBottles: 9, numExchange: 3}, want: 13},
		{name: "smoke 2", args: args{numBottles: 15, numExchange: 4}, want: 19},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numWaterBottles(test.args.numBottles, test.args.numExchange); got != test.want {
				t.Errorf("numWaterBottles() = %v, want = %v", got, test.want)
			}
		})
	}
}
