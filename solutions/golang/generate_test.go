package golang

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	type args struct {
		numRows int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "smoke 1", args: args{numRows: 5}, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{name: "smoke 2", args: args{numRows: 1}, want: [][]int{{1}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := generate(test.args.numRows); !reflect.DeepEqual(got, test.want) {
				t.Errorf("generate() = %v, want = %v", got, test.want)
			}
		})
	}
}
