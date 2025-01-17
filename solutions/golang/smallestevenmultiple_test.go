package golang

import "testing"

func TestSmallestEvenMultiple(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 5}, want: 10},
		{name: "smoke 2", args: args{n: 6}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := smallestEvenMultiple(test.args.n); got != test.want {
				t.Errorf("smallestEvenMultiple() = %v, want = %v", got, test.want)
			}
		})
	}
}
