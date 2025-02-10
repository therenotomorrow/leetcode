package golang

import "testing"

func TestClearDigits(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "abc"}, want: "abc"},
		{name: "smoke 2", args: args{s: "cb34"}, want: ""},
		{name: "test 267: wrong answer", args: args{s: "a8f"}, want: "f"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := clearDigits(test.args.s); got != test.want {
				t.Errorf("clearDigits() = %v, want = %v", got, test.want)
			}
		})
	}
}
