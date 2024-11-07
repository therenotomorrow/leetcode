package golang

import (
	"testing"
)

func TestLargestCombination(t *testing.T) {
	t.Parallel()

	type args struct {
		candidates []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{candidates: []int{16, 17, 71, 62, 12, 24, 14}}, want: 4},
		{name: "smoke 2", args: args{candidates: []int{8, 8}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestCombination(test.args.candidates); got != test.want {
				t.Errorf("largestCombination() = %v, want = %v", got, test.want)
			}
		})
	}
}
