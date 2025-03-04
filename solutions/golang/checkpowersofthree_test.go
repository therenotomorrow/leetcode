package golang

import "testing"

func TestCheckPowersOfThree(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 12}, want: true},
		{name: "smoke 2", args: args{n: 91}, want: true},
		{name: "smoke 3", args: args{n: 21}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := checkPowersOfThree(test.args.n); got != test.want {
				t.Errorf("checkPowersOfThree() = %v, want = %v", got, test.want)
			}
		})
	}
}
