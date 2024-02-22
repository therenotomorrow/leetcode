package golang

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 16}, want: true},
		{name: "smoke 2", args: args{n: 5}, want: false},
		{name: "smoke 3", args: args{n: 1}, want: true},
		{name: "test 1045: wrong answer", args: args{n: 0}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want = %v", got, tt.want)
			}
		})
	}
}
