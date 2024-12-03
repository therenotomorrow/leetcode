package golang

import (
	"testing"
)

func TestAddSpaces(t *testing.T) {
	t.Parallel()

	type args struct {
		s      string
		spaces []int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{s: "LeetcodeHelpsMeLearn", spaces: []int{8, 13, 15}},
			want: "Leetcode Helps Me Learn",
		},
		{
			name: "smoke 2",
			args: args{s: "icodeinpython", spaces: []int{1, 5, 7, 9}},
			want: "i code in py thon",
		},
		{
			name: "smoke 3",
			args: args{s: "spacing", spaces: []int{0, 1, 2, 3, 4, 5, 6}},
			want: " s p a c i n g",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := addSpaces(test.args.s, test.args.spaces); got != test.want {
				t.Errorf("addSpaces() = %v, want = %v", got, test.want)
			}
		})
	}
}
