package golang

import "testing"

func TestFindMiddleIndex(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 3, -1, 8, 4}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, -1, 4}}, want: 2},
		{name: "smoke 3", args: args{nums: []int{2, 5}}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want = %v", got, tt.want)
			}
		})
	}
}
