package golang

import "testing"

func TestLargestOddNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{num: "52"}, want: "5"},
		{name: "smoke 2", args: args{num: "4206"}, want: ""},
		{name: "smoke 3", args: args{num: "35427"}, want: "35427"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestOddNumber(test.args.num); got != test.want {
				t.Errorf("largestOddNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
