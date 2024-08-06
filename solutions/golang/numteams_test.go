package golang

import "testing"

func TestNumTeams(t *testing.T) {
	t.Parallel()

	type args struct {
		rating []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{rating: []int{2, 5, 3, 4, 1}}, want: 3},
		{name: "smoke 2", args: args{rating: []int{2, 1, 3}}, want: 0},
		{name: "smoke 3", args: args{rating: []int{1, 2, 3, 4}}, want: 4},
		{name: "test 13: wrong answer", args: args{rating: []int{4, 7, 9, 5, 10, 8, 2, 1, 6}}, want: 24},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numTeams(test.args.rating); got != test.want {
				t.Errorf("numTeams() = %v, want = %v", got, test.want)
			}
		})
	}
}
