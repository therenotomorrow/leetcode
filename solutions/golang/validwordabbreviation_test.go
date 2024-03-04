package golang

import "testing"

func TestValidWordAbbreviation(t *testing.T) {
	type args struct {
		word string
		abbr string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{word: "internationalization", abbr: "i12iz4n"}, want: true},
		{name: "smoke 2", args: args{word: "apple", abbr: "a2e"}, want: false},
		{name: "test 2: wrong answer", args: args{word: "internationalization", abbr: "i5a11o1"}, want: true},
		{name: "test 10: wrong answer", args: args{word: "a", abbr: "2"}, want: false},
		{name: "test 12: wrong answer", args: args{word: "a", abbr: "01"}, want: false},
		{name: "test 12: runtime", args: args{word: "hi", abbr: "2i"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validWordAbbreviation(tt.args.word, tt.args.abbr); got != tt.want {
				t.Errorf("validWordAbbreviation() = %v, want = %v", got, tt.want)
			}
		})
	}
}
