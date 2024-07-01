package golang

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Parallel()

	type args struct {
		s []byte
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "smoke 1", args: args{[]byte{'h', 'e', 'l', 'l', 'o'}}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{name: "smoke 2", args: args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			reverseString(test.args.s)

			if !reflect.DeepEqual(test.args.s, test.want) {
				t.Errorf("reverseString() = %v, want = %v", test.args.s, test.want)
			}
		})
	}
}
