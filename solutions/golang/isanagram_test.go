package golang

import "testing"

func TestIsAnagram(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "anagram", t: "nagaram"}, want: true},
		{name: "smoke 2", args: args{s: "rat", t: "car"}, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isAnagram(test.args.s, test.args.t); got != test.want {
				t.Errorf("isAnagram() = %v, want = %v", got, test.want)
			}
		})
	}
}
