package golang

import "testing"

func TestConvertToTitle(t *testing.T) {
	t.Parallel()

	type args struct {
		columnNumber int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{columnNumber: 1}, want: "A"},
		{name: "smoke 2", args: args{columnNumber: 28}, want: "AB"},
		{name: "smoke 3", args: args{columnNumber: 701}, want: "ZY"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := convertToTitle(test.args.columnNumber); got != test.want {
				t.Errorf("convertToTitle() = %v, want = %v", got, test.want)
			}
		})
	}
}
