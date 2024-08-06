package golang

import "testing"

func TestProbabilityOfHeads(t *testing.T) {
	t.Parallel()

	type args struct {
		prob   []float64
		target int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "smoke 1", args: args{prob: []float64{0.4}, target: 1}, want: 0.4},
		{name: "smoke 2", args: args{prob: []float64{0.5, 0.5, 0.5, 0.5, 0.5}, target: 0}, want: 0.03125},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := probabilityOfHeads(test.args.prob, test.args.target); got != test.want {
				t.Errorf("probabilityOfHeads() = %v, want = %v", got, test.want)
			}
		})
	}
}
