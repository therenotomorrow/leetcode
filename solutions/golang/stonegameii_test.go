package golang

import "testing"

func TestStoneGameII(t *testing.T) {
	t.Parallel()

	type args struct {
		piles []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{piles: []int{2, 7, 9, 4, 4}}, want: 10},
		{name: "smoke 2", args: args{piles: []int{1, 2, 3, 4, 5, 100}}, want: 104},
		{name: "test 23: wrong answer", args: args{piles: []int{77, 12, 64, 35, 28, 4, 87, 21, 20}}, want: 217},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := stoneGameII(test.args.piles); got != test.want {
				t.Errorf("stoneGameII() = %v, want = %v", got, test.want)
			}
		})
	}
}
