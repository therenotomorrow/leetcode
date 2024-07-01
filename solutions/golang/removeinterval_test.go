package golang

import (
	"reflect"
	"testing"
)

func TestRemoveInterval(t *testing.T) {
	t.Parallel()

	type args struct {
		intervals   [][]int
		toBeRemoved []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{intervals: [][]int{{0, 2}, {3, 4}, {5, 7}}, toBeRemoved: []int{1, 6}},
			want: [][]int{{0, 1}, {6, 7}},
		},
		{
			name: "smoke 2",
			args: args{intervals: [][]int{{0, 5}}, toBeRemoved: []int{2, 3}},
			want: [][]int{{0, 2}, {3, 5}},
		},
		{
			name: "smoke 3",
			args: args{
				intervals:   [][]int{{-5, -4}, {-3, -2}, {1, 2}, {3, 5}, {8, 9}},
				toBeRemoved: []int{-1, 4},
			},
			want: [][]int{{-5, -4}, {-3, -2}, {4, 5}, {8, 9}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := removeInterval(test.args.intervals, test.args.toBeRemoved); !reflect.DeepEqual(got, test.want) {
				t.Errorf("removeInterval() = %v, want = %v", got, test.want)
			}
		})
	}
}
