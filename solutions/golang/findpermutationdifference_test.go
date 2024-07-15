package golang

import "testing"

func TestFindPermutationDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{"abc", "bac"}, want: 2},
		{name: "smoke 2", args: args{"abcde", "edbac"}, want: 12},
		{name: "test 26: wrong answer", args: args{"rwohu", "rwuoh"}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findPermutationDifference(test.args.s, test.args.t); got != test.want {
				t.Errorf("findPermutationDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
