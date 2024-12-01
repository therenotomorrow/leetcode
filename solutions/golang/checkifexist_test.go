package golang

import (
	"testing"
)

func TestCheckIfExist(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{arr: []int{10, 2, 5, 3}}, want: true},
		{name: "smoke 2", args: args{arr: []int{3, 1, 7, 11}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := checkIfExist(test.args.arr); got != test.want {
				t.Errorf("checkIfExist() = %v, want = %v", got, test.want)
			}
		})
	}
}
