package golang

import "testing"

func TestRemoveDigit(t *testing.T) {
	t.Parallel()

	type args struct {
		number string
		digit  byte
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{number: "123", digit: '3'}, want: "12"},
		{name: "smoke 2", args: args{number: "1231", digit: '1'}, want: "231"},
		{name: "smoke 3", args: args{number: "551", digit: '5'}, want: "51"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := removeDigit(test.args.number, test.args.digit); got != test.want {
				t.Errorf("removeDigit() = %v, want = %v", got, test.want)
			}
		})
	}
}
