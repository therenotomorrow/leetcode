package golang

import "testing"

func TestTupleSameProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 3, 4, 6}}, want: 8},
		{name: "smoke 2", args: args{nums: []int{1, 2, 4, 5, 10}}, want: 16},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := tupleSameProduct(test.args.nums); got != test.want {
				t.Errorf("tupleSameProduct() = %v, want = %v", got, test.want)
			}
		})
	}
}
