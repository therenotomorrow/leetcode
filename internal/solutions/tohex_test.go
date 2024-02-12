package solutions

import "testing"

func TestToHex(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{num: 26}, want: "1a"},
		{name: "smoke 2", args: args{num: -1}, want: "ffffffff"},
		{name: "test 99: wrong answer", args: args{num: 0}, want: "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want = %v", got, tt.want)
			}
		})
	}
}
