package golang

import (
	"reflect"
	"testing"
)

func TestQueryResults(t *testing.T) {
	t.Parallel()

	type args struct {
		limit   int
		queries [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{limit: 4, queries: [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}},
			want: []int{1, 2, 2, 3},
		},
		{
			name: "smoke 2",
			args: args{limit: 4, queries: [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}},
			want: []int{1, 2, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := queryResults(test.args.limit, test.args.queries); !reflect.DeepEqual(got, test.want) {
				t.Errorf("queryResults() = %v, want = %v", got, test.want)
			}
		})
	}
}
