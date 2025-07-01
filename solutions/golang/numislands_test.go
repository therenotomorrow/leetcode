package golang

import "testing"

func TestNumIslands(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]byte
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{grid: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}}, want: 1},
		{name: "smoke 2", args: args{grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numIslands(test.args.grid); got != test.want {
				t.Errorf("numIslands() = %v, want = %v", got, test.want)
			}
		})
	}
}
