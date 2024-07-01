package golang

import "testing"

func TestMaxDepth(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "(1+(2*3)+((8)/4))+1"}, want: 3},
		{name: "smoke 2", args: args{s: "(1)+((2))+(((3)))"}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxDepth(test.args.s); got != test.want {
				t.Errorf("maxDepth() = %v, want = %v", got, test.want)
			}
		})
	}
}
