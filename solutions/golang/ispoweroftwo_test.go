package golang

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 1}, want: true},
		{name: "smoke 2", args: args{n: 16}, want: true},
		{name: "smoke 3", args: args{n: 3}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPowerOfTwo(test.args.n); got != test.want {
				t.Errorf("isPowerOfTwo() = %v, want = %v", got, test.want)
			}
		})
	}
}
