package golang

import "testing"

func TestGetHappyString(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{n: 1, k: 3}, want: "c"},
		{name: "smoke 2", args: args{n: 1, k: 4}, want: ""},
		{name: "smoke 3", args: args{n: 3, k: 9}, want: "cab"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getHappyString(test.args.n, test.args.k); got != test.want {
				t.Errorf("getHappyString() = %v, want = %v", got, test.want)
			}
		})
	}
}
