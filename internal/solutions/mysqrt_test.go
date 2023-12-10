package solutions

import "testing"

func TestMySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{x: 4}, want: 2},
		{name: "smoke 2", args: args{x: 8}, want: 2},
		{name: "test 1015: wrong answer", args: args{x: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want = %v", got, tt.want)
			}
		})
	}
}
