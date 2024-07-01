package golang

import "testing"

func TestReverseWords(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "Let's take LeetCode contest"}, want: "s'teL ekat edoCteeL tsetnoc"},
		{name: "smoke 2", args: args{s: "God Ding"}, want: "doG gniD"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reverseWords(test.args.s); got != test.want {
				t.Errorf("reverseWords() = %v, want = %v", got, test.want)
			}
		})
	}
}
