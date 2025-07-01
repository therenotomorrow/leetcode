package golang

import "testing"

func TestPossibleStringCount(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{word: "abbcccc"}, want: 5},
		{name: "smoke 2", args: args{word: "abcd"}, want: 1},
		{name: "smoke 3", args: args{word: "aaaa"}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := possibleStringCount(test.args.word); got != test.want {
				t.Errorf("possibleStringCount() = %v, want = %v", got, test.want)
			}
		})
	}
}
