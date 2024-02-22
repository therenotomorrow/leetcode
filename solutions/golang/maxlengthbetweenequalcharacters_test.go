package golang

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aa"}, want: 0},
		{name: "smoke 2", args: args{s: "abca"}, want: 2},
		{name: "smoke 3", args: args{s: "cbzxy"}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLengthBetweenEqualCharacters(tt.args.s); got != tt.want {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want = %v", got, tt.want)
			}
		})
	}
}
