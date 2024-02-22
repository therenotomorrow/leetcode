package golang

import "testing"

func TestUniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{m: 3, n: 7}, want: 28},
		{name: "smoke 2", args: args{m: 3, n: 2}, want: 3},
		{name: "test 37: wrong answer", args: args{m: 23, n: 12}, want: 193536720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want = %v", got, tt.want)
			}
		})
	}
}
