package golang

import "testing"

func TestMaximumOddBinaryNumber(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "010"}, want: "001"},
		{name: "smoke 2", args: args{s: "0101"}, want: "1001"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOddBinaryNumber(tt.args.s); got != tt.want {
				t.Errorf("maximumOddBinaryNumber() = %v, want = %v", got, tt.want)
			}
		})
	}
}
