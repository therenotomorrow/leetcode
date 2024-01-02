package decodeString

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "3[a]2[bc]"}, want: "aaabcbc"},
		{args: args{s: "3[a2[c]]"}, want: "accaccacc"},
		{args: args{s: "2[abc]3[cd]ef"}, want: "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
