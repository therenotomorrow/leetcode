package golang

import (
	"testing"
)

func TestCanMakeSubsequence(t *testing.T) {
	t.Parallel()

	type args struct {
		str1 string
		str2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{str1: "abc", str2: "ad"}, want: true},
		{name: "smoke 2", args: args{str1: "zc", str2: "ad"}, want: true},
		{name: "smoke 3", args: args{str1: "ab", str2: "d"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canMakeSubsequence(test.args.str1, test.args.str2); got != test.want {
				t.Errorf("canMakeSubsequence() = %v, want = %v", got, test.want)
			}
		})
	}
}
