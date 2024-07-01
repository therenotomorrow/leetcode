package golang

import "testing"

func TestPoorPigs(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := poorPigs(test.args.buckets, test.args.minutesToDie, test.args.minutesToTest); got != test.want {
				t.Errorf("poorPigs() = %v, want = %v", got, test.want)
			}
		})
	}
}
