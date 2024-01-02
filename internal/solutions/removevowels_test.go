package solutions

import "testing"

func TestRemoveVowels(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "leetcodeisacommunityforcoders"}, want: "ltcdscmmntyfrcdrs"},
		{name: "smoke 2", args: args{s: "aeiou"}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeVowels(tt.args.s); got != tt.want {
				t.Errorf("removeVowels() = %v, want = %v", got, tt.want)
			}
		})
	}
}
