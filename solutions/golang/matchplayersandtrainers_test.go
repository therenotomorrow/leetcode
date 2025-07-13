package golang

import "testing"

func TestMatchPlayersAndTrainers(t *testing.T) {
	t.Parallel()

	type args struct {
		players  []int
		trainers []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{players: []int{4, 7, 9}, trainers: []int{8, 2, 5, 8}}, want: 2},
		{name: "smoke 2", args: args{players: []int{1, 1, 1}, trainers: []int{10}}, want: 1},
		{name: "test 8: runtime", args: args{players: []int{4, 2}, trainers: []int{4, 4, 3}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := matchPlayersAndTrainers(test.args.players, test.args.trainers); got != test.want {
				t.Errorf("matchPlayersAndTrainers() = %v, want = %v", got, test.want)
			}
		})
	}
}
