package golang

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "smoke 2",
			args: args{intervals: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "test 118: wrong answer",
			args: args{intervals: [][]int{{1, 4}, {2, 3}}},
			want: [][]int{{1, 4}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := merge(test.args.intervals); !reflect.DeepEqual(got, test.want) {
				t.Errorf("merge() = %v, want = %v", got, test.want)
			}
		})
	}
}
