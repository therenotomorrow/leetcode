package numOfArrays

import "testing"

func Test_numOfArrays(t *testing.T) {
	type args struct {
		n int
		m int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 2, m: 3, k: 1}, want: 6},
		{args: args{n: 5, m: 2, k: 3}, want: 0},
		{args: args{n: 9, m: 1, k: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfArrays(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("numOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
