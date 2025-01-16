package golang

import "testing"

func TestIsValid(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{word: "234Adas"}, want: true},
		{name: "smoke 2", args: args{word: "b3"}, want: false},
		{name: "smoke 3", args: args{word: "a3$e"}, want: false},
		{name: "test 675: wrong answer", args: args{word: "UuE6"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isValid(test.args.word); got != test.want {
				t.Errorf("isValid() = %v, want = %v", got, test.want)
			}
		})
	}
}
