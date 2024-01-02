package reverseString

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{args: args{[]byte{'h', 'e', 'l', 'l', 'o'}}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{args: args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		reverseString(tt.args.s)
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
