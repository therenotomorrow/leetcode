package golang

import "testing"

func TestCountHomogenous(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abbcccaa"}, want: 13},
		{name: "smoke 2", args: args{s: "xy"}, want: 2},
		{name: "smoke 3", args: args{s: "zzzzz"}, want: 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countHomogenous(test.args.s); got != test.want {
				t.Errorf("countHomogenous() = %v, want = %v", got, test.want)
			}
		})
	}
}
