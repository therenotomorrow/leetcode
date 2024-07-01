package golang

import "testing"

func TestMinOperations1(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "0100"}, want: 1},
		{name: "smoke 2", args: args{s: "10"}, want: 0},
		{name: "smoke 3", args: args{s: "1111"}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations1(test.args.s); got != test.want {
				t.Errorf("minOperations1() = %v, want = %v", got, test.want)
			}
		})
	}
}
