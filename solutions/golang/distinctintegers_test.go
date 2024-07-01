package golang

import "testing"

func TestDistinctIntegers(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 5}, want: 4},
		{name: "smoke 2", args: args{n: 3}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := distinctIntegers(test.args.n); got != test.want {
				t.Errorf("distinctIntegers() = %v, want = %v", got, test.want)
			}
		})
	}
}
