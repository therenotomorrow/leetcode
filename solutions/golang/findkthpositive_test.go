package golang

import (
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
		k   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{2, 3, 4, 7, 11}, k: 5}, want: 9},
		{name: "smoke 2", args: args{arr: []int{1, 2, 3, 4}, k: 2}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findKthPositive(test.args.arr, test.args.k); got != test.want {
				t.Errorf("findKthPositive() = %v, want = %v", got, test.want)
			}
		})
	}
}
