package golang

import "testing"

func TestMaxLength(t *testing.T) {
	type args struct {
		arr []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []string{"un", "iq", "ue"}}, want: 4},
		{name: "smoke 2", args: args{arr: []string{"cha", "r", "act", "ers"}}, want: 6},
		{name: "smoke 3", args: args{arr: []string{"abcdefghijklmnopqrstuvwxyz"}}, want: 26},
		{name: "test 29: wrong answer", args: args{arr: []string{"a", "abc", "d", "de", "def"}}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.arr); got != tt.want {
				t.Errorf("maxLength() = %v, want = %v", got, tt.want)
			}
		})
	}
}
