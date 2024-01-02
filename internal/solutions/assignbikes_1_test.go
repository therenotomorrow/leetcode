package solutions

import (
	"reflect"
	"testing"
)

func TestAssignBikes1(t *testing.T) {
	type args struct {
		workers [][]int
		bikes   [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{
				workers: [][]int{{0, 0}, {2, 1}},
				bikes:   [][]int{{1, 2}, {3, 3}},
			},
			want: []int{1, 0},
		},
		{
			name: "smoke 2",
			args: args{
				workers: [][]int{{0, 0}, {1, 1}, {2, 0}},
				bikes:   [][]int{{1, 0}, {2, 2}, {2, 1}},
			},
			want: []int{0, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignBikes1(tt.args.workers, tt.args.bikes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignBikes1() = %v, want = %v", got, tt.want)
			}
		})
	}
}
