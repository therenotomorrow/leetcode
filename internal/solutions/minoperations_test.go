package solutions

import "testing"

func TestMinOperations(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "0100"}, want: 1},
		{name: "smoke 2", args: args{s: "10"}, want: 0},
		{name: "smoke 3", args: args{s: "1111"}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.s); got != tt.want {
				t.Errorf("minOperations() = %v, want = %v", got, tt.want)
			}
		})
	}
}
