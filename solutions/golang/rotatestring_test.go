package golang

import "testing"

func TestRotateString(t *testing.T) {
	t.Parallel()

	type args struct {
		s    string
		goal string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "abcde", goal: "cdeab"}, want: true},
		{name: "smoke 2", args: args{s: "abcde", goal: "abced"}, want: false},
		{name: "test 37: wrong answer", args: args{s: "defdefdefabcabc", goal: "defdefabcabcdef"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := rotateString(test.args.s, test.args.goal); got != test.want {
				t.Errorf("rotateString() = %v, want = %v", got, test.want)
			}
		})
	}
}
