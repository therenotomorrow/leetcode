package golang

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{num: 16}, want: true},
		{name: "smoke 2", args: args{num: 14}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want = %v", got, tt.want)
			}
		})
	}
}
