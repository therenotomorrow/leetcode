package golang

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{1, 3, 5}}, want: 4},
		{name: "smoke 2", args: args{arr: []int{2, 4, 6}}, want: 0},
		{name: "smoke 3", args: args{arr: []int{1, 2, 3, 4, 5, 6, 7}}, want: 16},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numOfSubarrays(test.args.arr); got != test.want {
				t.Errorf("numOfSubarrays() = %v, want = %v", got, test.want)
			}
		})
	}
}
