package golang

import "testing"

func TestNumTilePossibilities(t *testing.T) {
	t.Parallel()

	type args struct {
		tiles string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{tiles: "AAB"}, want: 8},
		{name: "smoke 2", args: args{tiles: "AAABBC"}, want: 188},
		{name: "smoke 3", args: args{tiles: "V"}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numTilePossibilities(test.args.tiles); got != test.want {
				t.Errorf("numTilePossibilities() = %v, want = %v", got, test.want)
			}
		})
	}
}
