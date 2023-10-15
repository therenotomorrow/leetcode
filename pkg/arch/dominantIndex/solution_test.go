package dominantIndex

import "testing"

func Test_dominantIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, 6, 1, 0}}, want: 1},
		{args: args{nums: []int{1, 2, 3, 4}}, want: -1},
		{args: args{nums: []int{0, 0, 3, 2}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
