package golang

import "testing"

func TestCountSubstrings(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abc"}, want: 3},
		{name: "smoke 2", args: args{s: "aaa"}, want: 6},
		{name: "test 6: wrong answer", args: args{s: "a"}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countSubstrings(test.args.s); got != test.want {
				t.Errorf("countSubstrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
