package solutions

import (
	"reflect"
	"testing"
)

func TestRemoveInterval(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInterval(tt.args.intervals, tt.args.toBeRemoved); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInterval() = %v, want = %v", got, tt.want)
			}
		})
	}
}
