package golang

import "testing"

func TestMinimumRecolors(t *testing.T) {
	t.Parallel()

	type args struct {
		blocks string
		k      int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{blocks: "WBBWWBBWBW", k: 7}, want: 3},
		{name: "smoke 2", args: args{blocks: "WBWBBBW", k: 2}, want: 0},
		{name: "test 111: wrong answer", args: args{blocks: "BWWWBB", k: 6}, want: 3},
		{
			name: "test 117: wrong answer",
			args: args{blocks: "BBBBBWWBBWBWBWWWBWBWBBBBWBBBBWBWBWBWBWWBWWBWBWWWWBBWWWWBWWWWBWBBWBBWBBWWW", k: 29},
			want: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumRecolors(test.args.blocks, test.args.k); got != test.want {
				t.Errorf("minimumRecolors() = %v, want = %v", got, test.want)
			}
		})
	}
}
