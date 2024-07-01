package golang

import "testing"

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
		k   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{5, 5, 4}, k: 1}, want: 1},
		{name: "smoke 2", args: args{arr: []int{4, 3, 1, 1, 3, 3, 2}, k: 3}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findLeastNumOfUniqueInts(test.args.arr, test.args.k); got != test.want {
				t.Errorf("findLeastNumOfUniqueInts() = %v, want = %v", got, test.want)
			}
		})
	}
}
