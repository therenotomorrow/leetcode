package golang

import "testing"

func TestMaxLength(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []string{"un", "iq", "ue"}}, want: 4},
		{name: "smoke 2", args: args{arr: []string{"cha", "r", "act", "ers"}}, want: 6},
		{name: "smoke 3", args: args{arr: []string{"abcdefghijklmnopqrstuvwxyz"}}, want: 26},
		{name: "test 29: wrong answer", args: args{arr: []string{"a", "abc", "d", "de", "def"}}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxLength(test.args.arr); got != test.want {
				t.Errorf("maxLength() = %v, want = %v", got, test.want)
			}
		})
	}
}
