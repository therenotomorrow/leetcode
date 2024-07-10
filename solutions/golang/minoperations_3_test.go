package golang

import "testing"

func TestMinOperations3(t *testing.T) {
	t.Parallel()

	type args struct {
		logs []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{logs: []string{"d1/", "d2/", "../", "d21/", "./"}}, want: 2},
		{name: "smoke 2", args: args{logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}}, want: 3},
		{name: "smoke 3", args: args{logs: []string{"d1/", "../", "../", "../"}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations3(test.args.logs); got != test.want {
				t.Errorf("minOperations3() = %v, want = %v", got, test.want)
			}
		})
	}
}
