package solutions

import "testing"

func TestIsPathCrossing(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPathCrossing(tt.args.path); got != tt.want {
				t.Errorf("isPathCrossing() = %v, want = %v", got, tt.want)
			}
		})
	}
}
