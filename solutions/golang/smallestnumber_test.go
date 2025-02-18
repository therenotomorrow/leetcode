package golang

import "testing"

func TestSmallestNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		pattern string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{pattern: "IIIDIDDD"}, want: "123549876"},
		{name: "smoke 2", args: args{pattern: "DDD"}, want: "4321"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := smallestNumber(test.args.pattern); got != test.want {
				t.Errorf("smallestNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
