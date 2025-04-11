package golang

import "testing"

func TestCountSymmetricIntegers(t *testing.T) {
	t.Parallel()

	type args struct {
		low  int
		high int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{low: 1, high: 100}, want: 9},
		{name: "smoke 2", args: args{low: 1200, high: 1230}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countSymmetricIntegers(test.args.low, test.args.high); got != test.want {
				t.Errorf("countSymmetricIntegers() = %v, want = %v", got, test.want)
			}
		})
	}
}
