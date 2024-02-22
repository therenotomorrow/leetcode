package golang

import "testing"

func TestMaxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}}, want: 4},
		{name: "smoke 2", args: args{arrays: [][]int{{1}, {1}}}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.arrays); got != tt.want {
				t.Errorf("maxDistance() = %v, want = %v", got, tt.want)
			}
		})
	}
}
