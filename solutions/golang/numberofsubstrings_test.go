package golang

import "testing"

func TestNumberOfSubstrings(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abcabc"}, want: 10},
		{name: "smoke 2", args: args{s: "aaacb"}, want: 3},
		{name: "smoke 3", args: args{s: "abc"}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfSubstrings(test.args.s); got != test.want {
				t.Errorf("numberOfSubstrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
