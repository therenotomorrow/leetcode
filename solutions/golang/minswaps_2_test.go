package golang

import "testing"

func TestMinSwaps2(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "][]["}, want: 1},
		{name: "smoke 2", args: args{s: "]]][[["}, want: 2},
		{name: "smoke 3", args: args{s: "[]"}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minSwaps2(test.args.s); got != test.want {
				t.Errorf("minSwaps2() = %v, want = %v", got, test.want)
			}
		})
	}
}
