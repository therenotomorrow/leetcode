package golang

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{n: 13}, want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "smoke 2", args: args{n: 2}, want: []int{1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lexicalOrder(test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("lexicalOrder() = %v, want = %v", got, test.want)
			}
		})
	}
}
