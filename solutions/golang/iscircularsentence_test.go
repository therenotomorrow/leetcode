package golang

import "testing"

func TestIsCircularSentence(t *testing.T) {
	t.Parallel()

	type args struct {
		sentence string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{sentence: "leetcode exercises sound delightful"}, want: true},
		{name: "smoke 2", args: args{sentence: "eetcode"}, want: true},
		{name: "smoke 3", args: args{sentence: "Leetcode is cool"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isCircularSentence(test.args.sentence); got != test.want {
				t.Errorf("isCircularSentence() = %v, want = %v", got, test.want)
			}
		})
	}
}
