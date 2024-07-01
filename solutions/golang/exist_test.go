package golang

import "testing"

func TestExist(t *testing.T) {
	t.Parallel()

	type args struct {
		board [][]byte
		word  string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCCED",
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "SEE",
			},
			want: true,
		},
		{
			name: "smoke 3",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCB",
			},
			want: false,
		},
		{
			name: "test 44: wrong answer",
			args: args{
				board: [][]byte{{'a', 'b'}, {'c', 'd'}},
				word:  "abcd",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := exist(test.args.board, test.args.word); got != test.want {
				t.Errorf("exist() = %v, want = %v", got, test.want)
			}
		})
	}
}
