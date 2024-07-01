package golang

import "testing"

func TestIsPathCrossing(t *testing.T) {
	t.Parallel()

	type args struct {
		path string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{path: "NES"}, want: false},
		{name: "smoke 2", args: args{path: "NESWW"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPathCrossing(test.args.path); got != test.want {
				t.Errorf("isPathCrossing() = %v, want = %v", got, test.want)
			}
		})
	}
}
