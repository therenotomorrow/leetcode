package golang

import "testing"

func TestCustomSortString(t *testing.T) {
	t.Parallel()

	type args struct {
		order string
		s     string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{order: "cba", s: "abcd"}, want: "cbad"},
		{name: "smoke 2", args: args{order: "cbafg", s: "abcd"}, want: "cbad"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := customSortString(test.args.order, test.args.s); got != test.want {
				t.Errorf("customSortString() = %v, want = %v", got, test.want)
			}
		})
	}
}
