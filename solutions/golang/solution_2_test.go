package golang

import (
	"reflect"
	"testing"
)

func TestSolution2(t *testing.T) {
	t.Parallel()

	type args struct {
		graph [][]int
		n     int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{graph: [][]int{{1, 1, 0}, {0, 1, 0}, {1, 1, 1}}, n: 3}, want: 1},
		{name: "smoke 2", args: args{graph: [][]int{{1, 0, 1}, {1, 1, 0}, {0, 1, 1}}, n: 3}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			knows := func(a int, b int) bool {
				return test.args.graph[a][b] == 1
			}

			got := solution2(knows)(test.args.n)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("solution2() = %v, want = %v", got, test.want)
			}
		})
	}
}
