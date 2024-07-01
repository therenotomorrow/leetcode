package golang

import (
	"reflect"
	"testing"
)

func TestImageSmoother(t *testing.T) {
	t.Parallel()

	type args struct {
		img [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{img: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			name: "smoke 2",
			args: args{img: [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}},
			want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := imageSmoother(test.args.img); !reflect.DeepEqual(got, test.want) {
				t.Errorf("imageSmoother() = %v, want = %v", got, test.want)
			}
		})
	}
}
