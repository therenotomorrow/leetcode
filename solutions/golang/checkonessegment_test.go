package golang

import (
	"testing"
)

func TestCheckOnesSegment(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "1001"}, want: false},
		{name: "smoke 2", args: args{s: "110"}, want: true},
		{name: "test 129: wrong answer", args: args{s: "1"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := checkOnesSegment(test.args.s); got != test.want {
				t.Errorf("checkOnesSegment() = %v, want = %v", got, test.want)
			}
		})
	}
}
