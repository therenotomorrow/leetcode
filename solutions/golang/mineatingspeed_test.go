package golang

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	t.Parallel()

	type args struct {
		piles []int
		h     int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{piles: []int{3, 6, 7, 11}, h: 8}, want: 4},
		{name: "smoke 2", args: args{piles: []int{30, 11, 23, 4, 20}, h: 5}, want: 30},
		{name: "smoke 3", args: args{piles: []int{30, 11, 23, 4, 20}, h: 6}, want: 23},
		{name: "test 10: wrong answer", args: args{piles: []int{312884470}, h: 312884469}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minEatingSpeed(test.args.piles, test.args.h); got != test.want {
				t.Errorf("minEatingSpeed() = %v, want = %v", got, test.want)
			}
		})
	}
}
