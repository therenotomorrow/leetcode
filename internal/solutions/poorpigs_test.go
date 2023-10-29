package solutions

import "testing"

func TestPoorPigs(t *testing.T) {
	type args struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{buckets: 4, minutesToDie: 15, minutesToTest: 15}, want: 2},
		{name: "smoke 2", args: args{buckets: 4, minutesToDie: 15, minutesToTest: 30}, want: 2},
		{name: "test 11: wrong answer", args: args{buckets: 8, minutesToDie: 15, minutesToTest: 40}, want: 2},
		{name: "test 17: wrong answer", args: args{buckets: 125, minutesToDie: 1, minutesToTest: 4}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poorPigs(tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest); got != tt.want {
				t.Errorf("poorPigs() = %v, want = %v", got, tt.want)
			}
		})
	}
}
