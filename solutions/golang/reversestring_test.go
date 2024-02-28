package golang

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
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

	for _, tt := range tests {
		reverseString(tt.args.s)
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want = %v", tt.args.s, tt.want)
			}
		})
	}
}
