package golang

import (
	"reflect"
	"testing"
)

func TestFindWinners(t *testing.T) {
	t.Parallel()

	type args struct {
		matches [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{
				matches: [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
			},
			want: [][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
		{
			name: "smoke 2",
			args: args{
				matches: [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}},
			},
			want: [][]int{{1, 2, 5, 6}, {}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findWinners(test.args.matches); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findWinners() = %v, want = %v", got, test.want)
			}
		})
	}
}
