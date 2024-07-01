package golang

import "testing"

func TestNumRescueBoats(t *testing.T) {
	t.Parallel()

	type args struct {
		people []int
		limit  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{people: []int{1, 2}, limit: 3}, want: 1},
		{name: "smoke 2", args: args{people: []int{3, 2, 2, 1}, limit: 3}, want: 3},
		{name: "smoke 3", args: args{people: []int{3, 5, 3, 4}, limit: 5}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numRescueBoats(test.args.people, test.args.limit); got != test.want {
				t.Errorf("numRescueBoats() = %v, want = %v", got, test.want)
			}
		})
	}
}
