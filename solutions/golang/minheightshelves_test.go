package golang

import "testing"

func TestMinHeightShelves(t *testing.T) {
	t.Parallel()

	type args struct {
		books      [][]int
		shelfWidth int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{books: [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, shelfWidth: 4},
			want: 6,
		},
		{
			name: "smoke 2",
			args: args{books: [][]int{{1, 3}, {2, 4}, {3, 2}}, shelfWidth: 6},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minHeightShelves(test.args.books, test.args.shelfWidth); got != test.want {
				t.Errorf("minHeightShelves() = %v, want = %v", got, test.want)
			}
		})
	}
}
