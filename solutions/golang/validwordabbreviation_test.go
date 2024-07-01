package golang

import "testing"

func TestValidWordAbbreviation(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := validWordAbbreviation(test.args.word, test.args.abbr); got != test.want {
				t.Errorf("validWordAbbreviation() = %v, want = %v", got, test.want)
			}
		})
	}
}
