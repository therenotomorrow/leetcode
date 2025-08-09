package golang

import "testing"

func TestSmallestCommonElement(t *testing.T) {
	t.Parallel()

	type args struct {
		mat [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{mat: [][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}},
			want: 5,
		},
		{
			name: "smoke 2",
			args: args{mat: [][]int{{1, 2, 3}, {2, 3, 4}, {2, 3, 5}}},
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := smallestCommonElement(test.args.mat); got != test.want {
				t.Errorf("smallestCommonElement() = %v, want = %v", got, test.want)
			}
		})
	}
}
