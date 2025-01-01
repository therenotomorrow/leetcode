package golang

import "testing"

func TestStringShift(t *testing.T) {
	t.Parallel()

	type args struct {
		s     string
		shift [][]int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{s: "abc", shift: [][]int{{0, 1}, {1, 2}}},
			want: "cab",
		},
		{
			name: "smoke 2",
			args: args{s: "abcdefg", shift: [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}},
			want: "efgabcd",
		},
		{
			name: "test 4: wrong answer",
			args: args{s: "wpdhhcj", shift: [][]int{{0, 7}, {1, 7}, {1, 0}, {1, 3}, {0, 3}, {0, 6}, {1, 2}}},
			want: "hcjwpdh",
		},
		{
			name: "test 7: runtime",
			args: args{
				s:     "xqgwkiqpif",
				shift: [][]int{{1, 4}, {0, 7}, {0, 8}, {0, 7}, {0, 6}, {1, 3}, {0, 1}, {1, 7}, {0, 5}, {0, 6}},
			},
			want: "qpifxqgwki",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := stringShift(test.args.s, test.args.shift); got != test.want {
				t.Errorf("stringShift() = %v, want = %v", got, test.want)
			}
		})
	}
}
