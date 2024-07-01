package golang

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	t.Parallel()

	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{left: 5, right: 7}, want: 4},
		{name: "smoke 2", args: args{left: 0, right: 0}, want: 0},
		{name: "smoke 3", args: args{left: 1, right: 2147483647}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := rangeBitwiseAnd(test.args.left, test.args.right); got != test.want {
				t.Errorf("rangeBitwiseAnd() = %v, want = %v", got, test.want)
			}
		})
	}
}
