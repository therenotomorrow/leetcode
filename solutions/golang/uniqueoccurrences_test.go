package golang

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{[]int{1, 2, 2, 1, 1, 3}}, want: true},
		{name: "smoke 2", args: args{[]int{1, 2}}, want: false},
		{name: "smoke 3", args: args{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := uniqueOccurrences(test.args.arr); got != test.want {
				t.Errorf("uniqueOccurrences() = %v, want = %v", got, test.want)
			}
		})
	}
}
