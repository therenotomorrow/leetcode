package golang

import "testing"

func TestLengthAfterTransformations(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abcyy", t: 2}, want: 7},
		{name: "smoke 2", args: args{s: "azbk", t: 1}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lengthAfterTransformations(test.args.s, test.args.t); got != test.want {
				t.Errorf("lengthAfterTransformations() = %v, want = %v", got, test.want)
			}
		})
	}
}
