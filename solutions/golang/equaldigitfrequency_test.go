package golang

import "testing"

func TestEqualDigitFrequency(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{"1212"}, want: 5},
		{name: "smoke 2", args: args{"12321"}, want: 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := equalDigitFrequency(test.args.s); got != test.want {
				t.Errorf("equalDigitFrequency() = %v, want = %v", got, test.want)
			}
		})
	}
}
