package golang

import "testing"

func TestNumKLenSubstrNoRepeats(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "havefunonleetcode", k: 5}, want: 6},
		{name: "smoke 2", args: args{s: "home", k: 5}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numKLenSubstrNoRepeats(test.args.s, test.args.k); got != test.want {
				t.Errorf("numKLenSubstrNoRepeats() = %v, want = %v", got, test.want)
			}
		})
	}
}
