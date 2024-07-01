package golang

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		intervals   [][]int
		newInterval []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "smoke 2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := insert(test.args.intervals, test.args.newInterval); !reflect.DeepEqual(got, test.want) {
				t.Errorf("insert() = %v, want = %v", got, test.want)
			}
		})
	}
}
