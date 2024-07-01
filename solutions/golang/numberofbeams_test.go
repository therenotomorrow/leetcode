package golang

import "testing"

func TestNumberOfBeams(t *testing.T) {
	t.Parallel()

	type args struct {
		bank []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{bank: []string{"011001", "000000", "010100", "001000"}}, want: 8},
		{name: "smoke 2", args: args{bank: []string{"000", "111", "000"}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfBeams(test.args.bank); got != test.want {
				t.Errorf("numberOfBeams() = %v, want = %v", got, test.want)
			}
		})
	}
}
