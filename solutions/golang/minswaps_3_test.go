package golang

import "testing"

func TestMinSwaps3(t *testing.T) {
	t.Parallel()

	type args struct {
		data []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{data: []int{1, 0, 1, 0, 1}}, want: 1},
		{name: "smoke 2", args: args{data: []int{0, 0, 0, 1, 0}}, want: 0},
		{name: "smoke 3", args: args{data: []int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minSwaps3(test.args.data); got != test.want {
				t.Errorf("minSwaps3() = %v, want = %v", got, test.want)
			}
		})
	}
}
