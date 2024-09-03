package golang

import "testing"

func TestBinaryGap(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 22}, want: 2},
		{name: "smoke 2", args: args{n: 8}, want: 0},
		{name: "smoke 3", args: args{n: 5}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := binaryGap(test.args.n); got != test.want {
				t.Errorf("binaryGap() = %v, want = %v", got, test.want)
			}
		})
	}
}
