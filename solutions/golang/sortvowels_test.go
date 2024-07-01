package golang

import "testing"

func TestSortVowels(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "lEetcOde"}, want: "lEOtcede"},
		{name: "smoke 2", args: args{s: "lYmpH"}, want: "lYmpH"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sortVowels(test.args.s); got != test.want {
				t.Errorf("sortVowels() = %v, want = %v", got, test.want)
			}
		})
	}
}
