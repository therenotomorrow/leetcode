package golang

import "testing"

func TestMaximumOddBinaryNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "010"}, want: "001"},
		{name: "smoke 2", args: args{s: "0101"}, want: "1001"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumOddBinaryNumber(test.args.s); got != test.want {
				t.Errorf("maximumOddBinaryNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
