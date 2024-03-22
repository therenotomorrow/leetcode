package isPalindrome

import "testing"

func Test_isPalindrome3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "A man, a plan, a canal: Panama"}, want: true},
		{args: args{s: "race a car"}, want: false},
		{args: args{s: " "}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome3(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome3() = %v, want %v", got, tt.want)
			}
		})
	}
}
