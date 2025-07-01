package golang

import "testing"

func TestMaxDistance2(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "NWSE", k: 1}, want: 3},
		{name: "smoke 2", args: args{s: "NSWWEW", k: 3}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxDistance2(test.args.s, test.args.k); got != test.want {
				t.Errorf("maxDistance2() = %v, want = %v", got, test.want)
			}
		})
	}
}
