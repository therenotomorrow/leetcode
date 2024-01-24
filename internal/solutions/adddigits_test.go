package solutions

import "testing"

func TestAddDigits(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{num: 38}, want: 2},
		{name: "smoke 2", args: args{num: 0}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want = %v", got, tt.want)
			}
		})
	}
}
