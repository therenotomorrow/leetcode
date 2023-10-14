package peakIndexInMountainArray

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{arr: []int{0, 1, 0}}, want: 1},
		{args: args{arr: []int{0, 2, 1, 0}}, want: 1},
		{args: args{arr: []int{0, 10, 5, 2}}, want: 1},
		{args: args{arr: []int{3, 4, 5, 1}}, want: 2},
		{args: args{arr: []int{0, 3, 5, 12, 2}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
