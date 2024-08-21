package golang

import "testing"

func TestFindComplement(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{num: 5}, want: 2},
		{name: "smoke 2", args: args{num: 1}, want: 0},
		{name: "test 144: wrong answer", args: args{num: 2}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findComplement(test.args.num); got != test.want {
				t.Errorf("findComplement() = %v, want = %v", got, test.want)
			}
		})
	}
}
