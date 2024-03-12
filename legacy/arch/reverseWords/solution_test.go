package reverseWords

import "testing"

func Test_reverseWords2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "the sky is blue"}, want: "blue is sky the"},
		{args: args{s: "  hello world  "}, want: "world hello"},
		{args: args{s: "a good   example"}, want: "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords2(tt.args.s); got != tt.want {
				t.Errorf("reverseWords2() = %v, want %v", got, tt.want)
			}
		})
	}
}
