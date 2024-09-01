package golang

import (
	"reflect"
	"testing"
)

func TestConstruct2DArray(t *testing.T) {
	t.Parallel()

	type args struct {
		original []int
		m        int
		n        int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "smoke 1", args: args{original: []int{1, 2, 3, 4}, m: 2, n: 2}, want: [][]int{{1, 2}, {3, 4}}},
		{name: "smoke 2", args: args{original: []int{1, 2, 3}, m: 1, n: 3}, want: [][]int{{1, 2, 3}}},
		{name: "smoke 3", args: args{original: []int{1, 2}, m: 1, n: 1}, want: [][]int{}},
		{name: "test 3: runtime", args: args{original: []int{3}, m: 1, n: 2}, want: [][]int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := construct2DArray(test.args.original, test.args.m, test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("construct2DArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
