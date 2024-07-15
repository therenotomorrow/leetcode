package golang

import "testing"

func TestBalancedStringSplit(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "RLRRLLRLRL"}, want: 4},
		{name: "smoke 2", args: args{s: "RLRRRLLRLL"}, want: 2},
		{name: "smoke 3", args: args{s: "LLLLRRRR"}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := balancedStringSplit(test.args.s); got != test.want {
				t.Errorf("balancedStringSplit() = %v, want = %v", got, test.want)
			}
		})
	}
}
