package golang

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	type args struct {
		encoded []int
		first   int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{encoded: []int{1, 2, 3}, first: 1}, want: []int{1, 0, 2, 1}},
		{name: "smoke 2", args: args{encoded: []int{6, 2, 7, 3}, first: 4}, want: []int{4, 2, 0, 7, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := decode(test.args.encoded, test.args.first); !reflect.DeepEqual(got, test.want) {
				t.Errorf("decode() = %v, want = %v", got, test.want)
			}
		})
	}
}
