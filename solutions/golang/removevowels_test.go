package golang

import "testing"

func TestRemoveVowels(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := removeVowels(test.args.s); got != test.want {
				t.Errorf("removeVowels() = %v, want = %v", got, test.want)
			}
		})
	}
}
