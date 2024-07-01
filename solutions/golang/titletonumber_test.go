package golang

import "testing"

func TestTitleToNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		columnTitle string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{columnTitle: "A"}, want: 1},
		{name: "smoke 2", args: args{columnTitle: "AB"}, want: 28},
		{name: "smoke 3", args: args{columnTitle: "ZY"}, want: 701},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := titleToNumber(test.args.columnTitle); got != test.want {
				t.Errorf("titleToNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
