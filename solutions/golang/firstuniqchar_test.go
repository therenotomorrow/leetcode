package golang

import "testing"

func TestFirstUniqChar(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "leetcode"}, want: 0},
		{name: "smoke 2", args: args{s: "loveleetcode"}, want: 2},
		{name: "smoke 3", args: args{s: "aabb"}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := firstUniqChar(test.args.s); got != test.want {
				t.Errorf("firstUniqChar() = %v, want = %v", got, test.want)
			}
		})
	}
}
