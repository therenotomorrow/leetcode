package golang

import "testing"

func TestCanArrange(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
		k   int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{arr: []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, k: 5}, want: true},
		{name: "smoke 2", args: args{arr: []int{1, 2, 3, 4, 5, 6}, k: 7}, want: true},
		{name: "smoke 3", args: args{arr: []int{1, 2, 3, 4, 5, 6}, k: 10}, want: false},
		{name: "test 44: wrong answer", args: args{arr: []int{-10, 10}, k: 2}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canArrange(test.args.arr, test.args.k); got != test.want {
				t.Errorf("canArrange() = %v, want = %v", got, test.want)
			}
		})
	}
}
