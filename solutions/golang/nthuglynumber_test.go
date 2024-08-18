package golang

import "testing"

func TestNthUglyNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 10}, want: 12},
		{name: "smoke 2", args: args{n: 1}, want: 1},
		{name: "test 8: wrong answer", args: args{n: 7}, want: 8},
		{name: "test 500: time limit", args: args{n: 1352}, want: 402653184},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := nthUglyNumber(test.args.n); got != test.want {
				t.Errorf("nthUglyNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
