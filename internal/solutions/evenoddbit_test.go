package solutions

import (
	"reflect"
	"testing"
)

func TestEvenOddBit(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{n: 17}, want: []int{2, 0}},
		{name: "smoke 2", args: args{n: 2}, want: []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOddBit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBit() = %v, want = %v", got, tt.want)
			}
		})
	}
}
