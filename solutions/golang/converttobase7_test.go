package golang

import "testing"

func TestConvertToBase7(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{num: 100}, want: "202"},
		{name: "smoke 2", args: args{num: -7}, want: "-10"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := convertToBase7(test.args.num); got != test.want {
				t.Errorf("convertToBase7() = %v, want = %v", got, test.want)
			}
		})
	}
}
