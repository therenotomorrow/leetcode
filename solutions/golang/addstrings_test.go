package golang

import "testing"

func TestAddStrings(t *testing.T) {
	t.Parallel()

	type args struct {
		num1 string
		num2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{num1: "11", num2: "123"}, want: "134"},
		{name: "smoke 2", args: args{num1: "456", num2: "77"}, want: "533"},
		{name: "smoke 3", args: args{num1: "0", num2: "0"}, want: "0"},
		{name: "test 291: wrong answer", args: args{num1: "1", num2: "9"}, want: "10"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := addStrings(test.args.num1, test.args.num2); got != test.want {
				t.Errorf("addStrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
