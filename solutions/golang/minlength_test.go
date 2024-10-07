package golang

import "testing"

func TestMinLength(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "ABFCACDB"}, want: 2},
		{name: "smoke 2", args: args{s: "ACBBD"}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minLength(test.args.s); got != test.want {
				t.Errorf("minLength() = %v, want = %v", got, test.want)
			}
		})
	}
}
