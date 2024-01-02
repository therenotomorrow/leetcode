package fullBloomFlowers

import (
	"reflect"
	"testing"
)

func Test_fullBloomFlowers(t *testing.T) {
	type args struct {
		flowers [][]int
		people  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{flowers: [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, people: []int{2, 3, 7, 11}}, want: []int{1, 2, 2, 2}},
		{args: args{flowers: [][]int{{1, 10}, {3, 3}}, people: []int{3, 3, 2}}, want: []int{2, 2, 1}},
		{args: args{flowers: [][]int{{19, 37}, {19, 38}, {19, 35}}, people: []int{6, 7, 21, 1, 13, 37, 5, 37, 46, 43}}, want: []int{0, 0, 3, 0, 0, 2, 0, 2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullBloomFlowers(tt.args.flowers, tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullBloomFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
