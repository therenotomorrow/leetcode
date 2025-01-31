package golang

import (
	"reflect"
	"testing"
)

func TestEventualSafeNodes(t *testing.T) {
	t.Parallel()

	type args struct {
		graph [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}}, want: []int{2, 4, 5, 6}},
		{name: "smoke 2", args: args{graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}}, want: []int{4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := eventualSafeNodes(test.args.graph); !reflect.DeepEqual(got, test.want) {
				t.Errorf("eventualSafeNodes() = %v, want = %v", got, test.want)
			}
		})
	}
}
