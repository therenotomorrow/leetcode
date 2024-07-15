package golang

import "testing"

func TestDefangIPaddr(t *testing.T) {
	t.Parallel()

	type args struct {
		address string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{address: "1.1.1.1"}, want: "1[.]1[.]1[.]1"},
		{name: "smoke 2", args: args{address: "255.100.50.0"}, want: "255[.]100[.]50[.]0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := defangIPaddr(test.args.address); got != test.want {
				t.Errorf("defangIPaddr() = %v, want = %v", got, test.want)
			}
		})
	}
}
