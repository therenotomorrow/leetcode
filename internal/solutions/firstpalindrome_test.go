package solutions

import "testing"

func TestFirstPalindrome(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want = %v", got, tt.want)
			}
		})
	}
}
