package golang

import "testing"

func TestFirstPalindrome(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{words: []string{"abc", "car", "ada", "racecar", "cool"}}, want: "ada"},
		{name: "smoke 2", args: args{words: []string{"notapalindrome", "racecar"}}, want: "racecar"},
		{name: "smoke 3", args: args{words: []string{"def", "ghi"}}, want: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := firstPalindrome(test.args.words); got != test.want {
				t.Errorf("firstPalindrome() = %v, want = %v", got, test.want)
			}
		})
	}
}
