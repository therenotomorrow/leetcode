package golang

import "testing"

func TestAreAlmostEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		s1 string
		s2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s1: "bank", s2: "kanb"}, want: true},
		{name: "smoke 2", args: args{s1: "attack", s2: "defend"}, want: false},
		{name: "smoke 3", args: args{s1: "kelb", s2: "kelb"}, want: true},
		{name: "test 120: wrong answer", args: args{s1: "aa", s2: "ac"}, want: false},
		{name: "test 128: wrong answer", args: args{s1: "qgqeg", s2: "gqgeq"}, want: false},
		{name: "test 130: wrong answer", args: args{s1: "aa", s2: "bb"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := areAlmostEqual(test.args.s1, test.args.s2); got != test.want {
				t.Errorf("areAlmostEqual() = %v, want = %v", got, test.want)
			}
		})
	}
}
