package golang

import "testing"

func TestHammingWeight(t *testing.T) {
	t.Parallel()

	type args struct {
		num uint32
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{num: 0o0000000000000000000000000001011}, want: 3},
		{name: "smoke 2", args: args{num: 0o0000000000000000000000010000000}, want: 1},
		{name: "smoke 3", args: args{num: 0o0000000000000000000010000100101}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := hammingWeight(test.args.num); got != test.want {
				t.Errorf("hammingWeight() = %v, want = %v", got, test.want)
			}
		})
	}
}
