package golang

import "testing"

func TestIsPalindrome1(t *testing.T) {
	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{x: 121}, want: true},
		{name: "smoke 2", args: args{x: -121}, want: false},
		{name: "smoke 3", args: args{x: 10}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome1(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome1() = %v, want = %v", got, tt.want)
			}
		})
	}
}
