package golang

import "testing"

func TestEqualSubstring(t *testing.T) {
	t.Parallel()

	type args struct {
		s       string
		t       string
		maxCost int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abcd", t: "bcdf", maxCost: 3}, want: 3},
		{name: "smoke 2", args: args{s: "abcd", t: "cdef", maxCost: 3}, want: 1},
		{name: "smoke 3", args: args{s: "abcd", t: "acde", maxCost: 0}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := equalSubstring(test.args.s, test.args.t, test.args.maxCost); got != test.want {
				t.Errorf("equalSubstring() = %v, want = %v", got, test.want)
			}
		})
	}
}
