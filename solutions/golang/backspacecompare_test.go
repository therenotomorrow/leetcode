package golang

import "testing"

func TestBackspaceCompare(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "ab#c", t: "ad#c"}, want: true},
		{name: "smoke 2", args: args{s: "ab##", t: "c#d#"}, want: true},
		{name: "smoke 3", args: args{s: "a#c", t: "b"}, want: false},
		{name: "test 1: runtime", args: args{s: "ab##", t: "c#d#"}, want: true},
		{name: "test 102: wrong answer", args: args{s: "y#fo##f", t: "y#f#o##f"}, want: true},
		{name: "test 106: wrong answer", args: args{s: "bxj##tw", t: "bxj###tw"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := backspaceCompare(test.args.s, test.args.t); got != test.want {
				t.Errorf("backspaceCompare() = %v, want = %v", got, test.want)
			}
		})
	}
}
