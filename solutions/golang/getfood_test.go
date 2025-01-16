package golang

import "testing"

func TestGetFood(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]byte
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{grid: [][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', '*', 'O', 'O', 'O', 'X'},
				{'X', 'O', 'O', '#', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X'},
			}},
			want: 3,
		},
		{
			name: "smoke 2",
			args: args{grid: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', '*', 'X', 'O', 'X'},
				{'X', 'O', 'X', '#', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			}},
			want: -1,
		},
		{
			name: "smoke 3",
			args: args{grid: [][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', '*', 'O', 'X', 'O', '#', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
			}},
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getFood(test.args.grid); got != test.want {
				t.Errorf("getFood() = %v, want = %v", got, test.want)
			}
		})
	}
}
