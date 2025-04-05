package golang

import "testing"

func TestShortestWay(t *testing.T) {
	t.Parallel()

	type args struct {
		source string
		target string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{source: "abc", target: "abcbc"}, want: 2},
		{name: "smoke 2", args: args{source: "abc", target: "acdbc"}, want: -1},
		{name: "smoke 3", args: args{source: "xyz", target: "xzyxz"}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := shortestWay(test.args.source, test.args.target); got != test.want {
				t.Errorf("shortestWay() = %v, want = %v", got, test.want)
			}
		})
	}
}
