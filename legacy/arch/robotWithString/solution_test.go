package robotWithString

import "testing"

func Test_robotWithString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "zza"}, want: "azz"},
		{args: args{s: "bac"}, want: "abc"},
		{args: args{s: "bdda"}, want: "addb"},
		{args: args{s: "bydizfve"}, want: "bdevfziy"},
		{args: args{s: "vzhofnpo"}, want: "fnohopzv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotWithString(tt.args.s); got != tt.want {
				t.Errorf("robotWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}
