package golang

import (
	"testing"
)

func TestMinChanges(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "1001"}, want: 2},
		{name: "smoke 2", args: args{s: "10"}, want: 1},
		{name: "smoke 3", args: args{s: "0000"}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minChanges(test.args.s); got != test.want {
				t.Errorf("minChanges() = %v, want = %v", got, test.want)
			}
		})
	}
}
