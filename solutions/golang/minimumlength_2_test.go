package golang

import "testing"

func TestMinimumLength2(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{"abaacbcbb"}, want: 5},
		{name: "smoke 2", args: args{"aa"}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumLength2(test.args.s); got != test.want {
				t.Errorf("minimumLength2() = %v, want = %v", got, test.want)
			}
		})
	}
}
