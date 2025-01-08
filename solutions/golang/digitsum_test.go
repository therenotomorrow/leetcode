package golang

import "testing"

func TestDigitSum(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "11111222223", k: 3}, want: "135"},
		{name: "smoke 2", args: args{s: "00000000", k: 3}, want: "000"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := digitSum(test.args.s, test.args.k); got != test.want {
				t.Errorf("digitSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
