package golang

import "testing"

func TestRobotWithString(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "zza"}, want: "azz"},
		{name: "smoke 2", args: args{s: "bac"}, want: "abc"},
		{name: "smoke 3", args: args{s: "bdda"}, want: "addb"},
		{name: "test 4: wrong answer", args: args{s: "bydizfve"}, want: "bdevfziy"},
		{name: "test 7: wrong answer", args: args{s: "vzhofnpo"}, want: "fnohopzv"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := robotWithString(test.args.s); got != test.want {
				t.Errorf("robotWithString() = %v, want = %v", got, test.want)
			}
		})
	}
}
