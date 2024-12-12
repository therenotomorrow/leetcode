package golang

import (
	"testing"
)

func TestPickGifts(t *testing.T) {
	t.Parallel()

	type args struct {
		gifts []int
		k     int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{gifts: []int{25, 64, 9, 4, 100}, k: 4}, want: 29},
		{name: "smoke 2", args: args{gifts: []int{1, 1, 1, 1}, k: 4}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := pickGifts(test.args.gifts, test.args.k); got != test.want {
				t.Errorf("pickGifts() = %v, want = %v", got, test.want)
			}
		})
	}
}
