package golang

import "testing"

func TestMaximumHappinessSum(t *testing.T) {
	t.Parallel()

	type args struct {
		happiness []int
		k         int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{happiness: []int{1, 2, 3}, k: 2}, want: 4},
		{name: "smoke 2", args: args{happiness: []int{1, 1, 1, 1}, k: 2}, want: 1},
		{name: "smoke 3", args: args{happiness: []int{2, 3, 4, 5}, k: 1}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumHappinessSum(test.args.happiness, test.args.k); got != test.want {
				t.Errorf("maximumHappinessSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
