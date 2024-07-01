package golang

import "testing"

func TestHalvesAreAlike(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "book"}, want: true},
		{name: "smoke 2", args: args{s: "textbook"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := halvesAreAlike(test.args.s); got != test.want {
				t.Errorf("halvesAreAlike() = %v, want = %v", got, test.want)
			}
		})
	}
}
