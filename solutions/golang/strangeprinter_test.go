package golang

import "testing"

func TestStrangePrinter(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aaabbb"}, want: 2},
		{name: "smoke 2", args: args{s: "aba"}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := strangePrinter(test.args.s); got != test.want {
				t.Errorf("strangePrinter() = %v, want = %v", got, test.want)
			}
		})
	}
}
