package golang

import "testing"

func TestToLowerCase(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "Hello"}, want: "hello"},
		{name: "smoke 2", args: args{s: "here"}, want: "here"},
		{name: "smoke 3", args: args{s: "LOVELY"}, want: "lovely"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := toLowerCase(test.args.s); got != test.want {
				t.Errorf("toLowerCase() = %v, want = %v", got, test.want)
			}
		})
	}
}
