package golang

import "testing"

func TestAreSentencesSimilar(t *testing.T) {
	t.Parallel()

	type args struct {
		sentence1 string
		sentence2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{sentence1: "My name is Haley", sentence2: "My Haley"}, want: true},
		{name: "smoke 2", args: args{sentence1: "of", sentence2: "A lot of words"}, want: false},
		{name: "smoke 3", args: args{sentence1: "Eating right now", sentence2: "Eating"}, want: true},
		{
			name: "test 122: wrong answer",
			args: args{sentence1: "eTUny i b R UFKQJ EZx JBJ Q xXz", sentence2: "eTUny i R EZx JBJ xXz"},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := areSentencesSimilar(test.args.sentence1, test.args.sentence2); got != test.want {
				t.Errorf("areSentencesSimilar() = %v, want = %v", got, test.want)
			}
		})
	}
}
