package golang

import (
	"reflect"
	"testing"
)

func TestMissingRolls(t *testing.T) {
	t.Parallel()

	type args struct {
		rolls []int
		mean  int
		n     int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{rolls: []int{3, 2, 4, 3}, mean: 4, n: 2}, want: []int{6, 6}},
		{name: "smoke 2", args: args{rolls: []int{1, 5, 6}, mean: 3, n: 4}, want: []int{3, 2, 2, 2}},
		{name: "smoke 3", args: args{rolls: []int{1, 2, 3, 4}, mean: 6, n: 4}, want: []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := missingRolls(test.args.rolls, test.args.mean, test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("missingRolls() = %v, want = %v", got, test.want)
			}
		})
	}
}
