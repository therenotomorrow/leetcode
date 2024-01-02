package largestSubmatrix

import "testing"

func TestLargestSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{matrix: [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}}, want: 4},
		{args: args{matrix: [][]int{{1, 0, 1, 0, 1}}}, want: 3},
		{args: args{matrix: [][]int{{1, 1, 0}, {1, 0, 1}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubmatrix(tt.args.matrix); got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
