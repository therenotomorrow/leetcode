package golang

import "testing"

func TestKthCharacter(t *testing.T) {
	t.Parallel()

	type args struct {
		k int
	}

	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "smoke 1", args: args{k: 5}, want: 'b'},
		{name: "smoke 2", args: args{k: 10}, want: 'c'},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := kthCharacter(test.args.k); got != test.want {
				t.Errorf("kthCharacter() = %v, want = %v", got, test.want)
			}
		})
	}
}
