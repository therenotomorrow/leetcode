package golang

import "testing"

func TestPivotInteger(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 8}, want: 6},
		{name: "smoke 2", args: args{n: 1}, want: 1},
		{name: "smoke 3", args: args{n: 4}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := pivotInteger(test.args.n); got != test.want {
				t.Errorf("pivotInteger() = %v, want = %v", got, test.want)
			}
		})
	}
}
