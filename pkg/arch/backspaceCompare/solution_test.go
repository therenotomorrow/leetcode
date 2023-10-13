package backspaceCompare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "ab#c", t: "ad#c"}, want: true},
		{args: args{s: "ab##", t: "c#d#"}, want: true},
		{args: args{s: "a#c", t: "b"}, want: false},
		{args: args{s: "y#fo##f", t: "y#f#o##f"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
