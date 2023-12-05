package solutions

import "testing"

func TestNumberOfMatches(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 7}, want: 6},
		{name: "smoke 2", args: args{n: 14}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfMatches(tt.args.n); got != tt.want {
				t.Errorf("numberOfMatches() = %v, want = %v", got, tt.want)
			}
		})
	}
}
