package golang

import "testing"

func TestSumOfMultiples(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 7}, want: 21},
		{name: "smoke 2", args: args{n: 10}, want: 40},
		{name: "smoke 3", args: args{n: 9}, want: 30},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sumOfMultiples(test.args.n); got != test.want {
				t.Errorf("sumOfMultiples() = %v, want = %v", got, test.want)
			}
		})
	}
}
