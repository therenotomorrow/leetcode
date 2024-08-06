package golang

import "testing"

func TestMinimumPushes(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{word: "abcde"}, want: 5},
		{name: "smoke 2", args: args{word: "xyzxyzxyzxyz"}, want: 12},
		{name: "smoke 3", args: args{word: "aabbccddeeffgghhiiiiii"}, want: 24},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumPushes(test.args.word); got != test.want {
				t.Errorf("minimumPushes() = %v, want = %v", got, test.want)
			}
		})
	}
}
