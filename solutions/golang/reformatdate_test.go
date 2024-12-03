package golang

import (
	"testing"
)

func TestReformatDate(t *testing.T) {
	t.Parallel()

	type args struct {
		date string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{date: "20th Oct 2052"}, want: "2052-10-20"},
		{name: "smoke 2", args: args{date: "6th Jun 1933"}, want: "1933-06-06"},
		{name: "smoke 3", args: args{date: "26th May 1960"}, want: "1960-05-26"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reformatDate(test.args.date); got != test.want {
				t.Errorf("reformatDate() = %v, want = %v", got, test.want)
			}
		})
	}
}
