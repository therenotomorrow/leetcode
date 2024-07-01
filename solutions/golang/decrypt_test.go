package golang

import (
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	t.Parallel()

	type args struct {
		code []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{code: []int{5, 7, 1, 4}, k: 3}, want: []int{12, 10, 16, 13}},
		{name: "smoke 2", args: args{code: []int{1, 2, 3, 4}, k: 0}, want: []int{0, 0, 0, 0}},
		{name: "smoke 3", args: args{code: []int{2, 4, 9, 3}, k: -2}, want: []int{12, 5, 6, 13}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := decrypt(test.args.code, test.args.k); !reflect.DeepEqual(got, test.want) {
				t.Errorf("decrypt() = %v, want = %v", got, test.want)
			}
		})
	}
}
