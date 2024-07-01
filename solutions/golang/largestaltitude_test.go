package golang

import "testing"

func TestLargestAltitude(t *testing.T) {
	t.Parallel()

	type args struct {
		gain []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{gain: []int{-5, 1, 5, 0, -7}}, want: 1},
		{name: "smoke 2", args: args{gain: []int{-4, -3, -2, -1, 4, 3, 2}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestAltitude(test.args.gain); got != test.want {
				t.Errorf("largestAltitude() = %v, want = %v", got, test.want)
			}
		})
	}
}
