package golang

import "testing"

func TestRemoveOccurrences(t *testing.T) {
	t.Parallel()

	type args struct {
		s    string
		part string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "daabcbaabcbc", part: "abc"}, want: "dab"},
		{name: "smoke 2", args: args{s: "axxxxyyyyb", part: "xy"}, want: "ab"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := removeOccurrences(test.args.s, test.args.part); got != test.want {
				t.Errorf("removeOccurrences() = %v, want = %v", got, test.want)
			}
		})
	}
}
