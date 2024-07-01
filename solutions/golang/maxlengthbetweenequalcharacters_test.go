package golang

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aa"}, want: 0},
		{name: "smoke 2", args: args{s: "abca"}, want: 2},
		{name: "smoke 3", args: args{s: "cbzxy"}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxLengthBetweenEqualCharacters(test.args.s); got != test.want {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want = %v", got, test.want)
			}
		})
	}
}
