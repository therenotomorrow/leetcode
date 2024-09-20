package golang

import (
	"reflect"
	"testing"
)

func TestXorQueries(t *testing.T) {
	t.Parallel()

	type args struct {
		arr     []int
		queries [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{arr: []int{1, 3, 4, 8}, queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}},
			want: []int{2, 7, 14, 8},
		},
		{
			name: "smoke 2",
			args: args{arr: []int{4, 8, 2, 10}, queries: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}},
			want: []int{8, 0, 4, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := xorQueries(test.args.arr, test.args.queries); !reflect.DeepEqual(got, test.want) {
				t.Errorf("xorQueries() = %v, want = %v", got, test.want)
			}
		})
	}
}
