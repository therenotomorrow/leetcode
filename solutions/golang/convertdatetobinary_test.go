package golang

import "testing"

func TestConvertDateToBinary(t *testing.T) {
	t.Parallel()

	type args struct {
		date string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{date: "2080-02-29"}, want: "100000100000-10-11101"},
		{name: "smoke 2", args: args{date: "1900-01-01"}, want: "11101101100-1-1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := convertDateToBinary(test.args.date); got != test.want {
				t.Errorf("convertDateToBinary() = %v, want = %v", got, test.want)
			}
		})
	}
}
