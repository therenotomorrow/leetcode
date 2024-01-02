package findInMountainArray

import "testing"

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		target      int
		mountainArr *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{target: 3, mountainArr: &MountainArray{data: []int{1, 2, 3, 4, 5, 3, 1}}}, want: 2},
		{args: args{target: 3, mountainArr: &MountainArray{data: []int{0, 1, 2, 4, 2, 1}}}, want: -1},
		{args: args{target: 1, mountainArr: &MountainArray{data: []int{1, 5, 2}}}, want: 0},
		{args: args{target: 2, mountainArr: &MountainArray{data: []int{1, 5, 2}}}, want: 2},
		{args: args{target: 5, mountainArr: &MountainArray{data: []int{1, 5, 2}}}, want: 1},
		{args: args{target: 1, mountainArr: &MountainArray{data: []int{0, 5, 3, 1}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
