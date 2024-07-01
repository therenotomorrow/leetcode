package golang

import "testing"

func TestLargestGoodInteger(t *testing.T) {
	t.Parallel()

	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{num: "6777133339"}, want: "777"},
		{name: "smoke 2", args: args{num: "2300019"}, want: "000"},
		{name: "smoke 3", args: args{num: "42352338"}, want: ""},
		{name: "test 71: wrong answer", args: args{num: "3200014888"}, want: "888"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestGoodInteger(test.args.num); got != test.want {
				t.Errorf("largestGoodInteger() = %v, want = %v", got, test.want)
			}
		})
	}
}
