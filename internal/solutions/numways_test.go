package solutions

import "testing"

func TestNumWays(t *testing.T) {
	type args struct {
		steps  int
		arrLen int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{steps: 3, arrLen: 2}, want: 4},
		{name: "smoke 2", args: args{steps: 2, arrLen: 4}, want: 2},
		{name: "smoke 3", args: args{steps: 4, arrLen: 2}, want: 8},
		{name: "test 7: wrong answer", args: args{steps: 4, arrLen: 3}, want: 9},
		{name: "test 9: time limit", args: args{steps: 13, arrLen: 11}, want: 41835},
		{name: "test 18: wrong answer", args: args{steps: 27, arrLen: 7}, want: 127784505},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.steps, tt.args.arrLen); got != tt.want {
				t.Errorf("numWays() = %v, want = %v", got, tt.want)
			}
		})
	}
}
