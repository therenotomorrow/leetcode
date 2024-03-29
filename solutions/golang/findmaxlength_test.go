package golang

import "testing"

func TestFindMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{0, 1, 0}}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want = %v", got, tt.want)
			}
		})
	}
}
