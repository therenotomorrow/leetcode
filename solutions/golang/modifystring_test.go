package golang

import (
	"testing"
)

func TestModifyString(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "?zs"}, want: "azs"},
		{name: "smoke 2", args: args{s: "ubv?w"}, want: "ubvaw"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := modifyString(test.args.s); got != test.want {
				t.Errorf("modifyString() = %v, want = %v", got, test.want)
			}
		})
	}
}
