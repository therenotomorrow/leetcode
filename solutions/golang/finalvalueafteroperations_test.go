package golang

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	t.Parallel()

	type args struct {
		operations []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{operations: []string{"--X", "X++", "X++"}}, want: 1},
		{name: "smoke 2", args: args{operations: []string{"++X", "++X", "X++"}}, want: 3},
		{name: "smoke 3", args: args{operations: []string{"X++", "++X", "--X", "X--"}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := finalValueAfterOperations(test.args.operations); got != test.want {
				t.Errorf("finalValueAfterOperations() = %v, want = %v", got, test.want)
			}
		})
	}
}
