package golang

import (
	"reflect"
	"testing"
)

func TestGetFinalState(t *testing.T) {
	t.Parallel()

	type args struct {
		nums       []int
		k          int
		multiplier int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2}, want: []int{8, 4, 6, 5, 6}},
		{name: "smoke 2", args: args{nums: []int{1, 2}, k: 3, multiplier: 4}, want: []int{16, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getFinalState(test.args.nums, test.args.k, test.args.multiplier); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getFinalState() = %v, want = %v", got, test.want)
			}
		})
	}
}
