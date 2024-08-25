package golang

import "testing"

func TestFractionAddition(t *testing.T) {
	t.Parallel()

	type args struct {
		expression string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{expression: "-1/2+1/2"}, want: "0/1"},
		{name: "smoke 2", args: args{expression: "-1/2+1/2+1/3"}, want: "1/3"},
		{name: "smoke 3", args: args{expression: "1/3-1/2"}, want: "-1/6"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := fractionAddition(test.args.expression); got != test.want {
				t.Errorf("fractionAddition() = %v, want = %v", got, test.want)
			}
		})
	}
}
