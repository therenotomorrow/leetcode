package golang

import "testing"

func TestThreeConsecutiveOdds(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{arr: []int{2, 6, 4, 1}}, want: false},
		{name: "smoke 2", args: args{arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := threeConsecutiveOdds(test.args.arr); got != test.want {
				t.Errorf("threeConsecutiveOdds() = %v, want = %v", got, test.want)
			}
		})
	}
}
