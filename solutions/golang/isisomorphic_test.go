package golang

import "testing"

func TestIsIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "egg", t: "add"}, want: true},
		{name: "smoke 2", args: args{s: "foo", t: "bar"}, want: false},
		{name: "smoke 3", args: args{s: "paper", t: "title"}, want: true},
		{name: "test 35: wrong answer", args: args{s: "badc", t: "baba"}, want: false},
		{name: "test 42: wrong answer", args: args{s: "bbbaaaba", t: "aaabbbba"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want = %v", got, tt.want)
			}
		})
	}
}
