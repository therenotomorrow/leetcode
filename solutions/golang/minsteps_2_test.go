package golang

import "testing"

func TestMinSteps2(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 3}, want: 3},
		{name: "smoke 2", args: args{n: 1}, want: 0},
		{name: "test 3: wrong answer", args: args{n: 4}, want: 4},
		{name: "test 7: wrong answer", args: args{n: 7}, want: 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minSteps2(test.args.n); got != test.want {
				t.Errorf("minSteps2() = %v, want = %v", got, test.want)
			}
		})
	}
}
