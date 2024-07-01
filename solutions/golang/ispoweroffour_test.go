package golang

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 16}, want: true},
		{name: "smoke 2", args: args{n: 5}, want: false},
		{name: "smoke 3", args: args{n: 1}, want: true},
		{name: "test 1045: wrong answer", args: args{n: 0}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPowerOfFour(test.args.n); got != test.want {
				t.Errorf("isPowerOfFour() = %v, want = %v", got, test.want)
			}
		})
	}
}
