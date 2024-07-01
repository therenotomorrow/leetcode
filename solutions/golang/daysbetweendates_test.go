package golang

import "testing"

func TestDaysBetweenDates(t *testing.T) {
	t.Parallel()

	type args struct {
		date1 string
		date2 string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{date1: "2019-06-29", date2: "2019-06-30"}, want: 1},
		{name: "smoke 2", args: args{date1: "2020-01-15", date2: "2019-12-31"}, want: 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := daysBetweenDates(test.args.date1, test.args.date2); got != test.want {
				t.Errorf("daysBetweenDates() = %v, want = %v", got, test.want)
			}
		})
	}
}
