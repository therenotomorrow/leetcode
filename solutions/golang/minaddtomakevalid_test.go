package golang

import "testing"

func TestMinAddToMakeValid(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "())"}, want: 1},
		{name: "smoke 2", args: args{s: "((("}, want: 3},
		{name: "test 44", args: args{s: "()))(("}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minAddToMakeValid(test.args.s); got != test.want {
				t.Errorf("minAddToMakeValid() = %v, want = %v", got, test.want)
			}
		})
	}
}
