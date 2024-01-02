package myAtoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "42"}, want: 42},
		{args: args{s: "   -42"}, want: -42},
		{args: args{s: "4193 with words"}, want: 4193},
		{args: args{s: "-91283472332"}, want: -2147483648},
		{args: args{s: ""}, want: 0},
		{args: args{s: "9223372036854775808"}, want: 2147483647},
		{args: args{s: "18446744073709551617"}, want: 2147483647},
		{args: args{s: "  0000000000012345678"}, want: 12345678},
		{args: args{s: "00000-42a1234"}, want: 0},
		{args: args{s: "000000000000000000"}, want: 0},
		{args: args{s: "-000000000000001"}, want: -1},
		{args: args{s: "-"}, want: 0},
		{args: args{s: "  -0 451"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
