package golang

import (
	"testing"
)

func TestMaximumPopulation(t *testing.T) {
	t.Parallel()

	type args struct {
		logs [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{logs: [][]int{{1993, 1999}, {2000, 2010}}}, want: 1993},
		{name: "smoke 2", args: args{logs: [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}}, want: 1960},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumPopulation(test.args.logs); got != test.want {
				t.Errorf("maximumPopulation() = %v, want = %v", got, test.want)
			}
		})
	}
}
