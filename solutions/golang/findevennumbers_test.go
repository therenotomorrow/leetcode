package golang

import (
	"reflect"
	"testing"
)

func TestFindEvenNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		digits []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{digits: []int{2, 1, 3, 0}},
			want: []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
		{name: "smoke 2", args: args{digits: []int{2, 2, 8, 8, 2}}, want: []int{222, 228, 282, 288, 822, 828, 882}},
		{name: "smoke 3", args: args{digits: []int{3, 7, 5}}, want: []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findEvenNumbers(test.args.digits); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findEvenNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
