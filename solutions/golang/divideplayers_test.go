package golang

import "testing"

func TestDividePlayers(t *testing.T) {
	t.Parallel()

	type args struct {
		skill []int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{skill: []int{3, 2, 5, 1, 3, 4}}, want: 22},
		{name: "smoke 2", args: args{skill: []int{3, 4}}, want: 12},
		{name: "smoke 3", args: args{skill: []int{1, 1, 2, 3}}, want: -1},
		{name: "test 80: wrong answer", args: args{skill: []int{2, 1, 5, 2}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := dividePlayers(test.args.skill); got != test.want {
				t.Errorf("dividePlayers() = %v, want = %v", got, test.want)
			}
		})
	}
}
