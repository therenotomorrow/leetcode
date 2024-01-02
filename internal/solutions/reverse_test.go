package solutions

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{x: 123}, want: 321},
		{name: "smoke 2", args: args{x: -123}, want: -321},
		{name: "smoke 3", args: args{x: 120}, want: 21},
		{name: "test 1036: wrong answer", args: args{x: 1534236469}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want = %v", got, tt.want)
			}
		})
	}
}
