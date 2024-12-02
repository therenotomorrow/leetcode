package golang

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	t.Parallel()

	type args struct {
		k int
		n int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "smoke 1", args: args{k: 3, n: 7}, want: [][]int{{1, 2, 4}}},
		{name: "smoke 2", args: args{k: 3, n: 9}, want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{name: "smoke 3", args: args{k: 4, n: 1}, want: [][]int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := combinationSum3(test.args.k, test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("combinationSum3() = %v, want = %v", got, test.want)
			}
		})
	}
}
