package golang

import "testing"

func TestCountSegments(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "Hello, my name is John"}, want: 5},
		{name: "smoke 2", args: args{s: "Hello"}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countSegments(test.args.s); got != test.want {
				t.Errorf("countSegments() = %v, want = %v", got, test.want)
			}
		})
	}
}
