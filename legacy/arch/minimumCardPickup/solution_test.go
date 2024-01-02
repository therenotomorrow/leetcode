package minimumCardPickup

import "testing"

func Test_minimumCardPickup(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{cards: []int{3, 4, 2, 3, 4, 7}}, want: 4},
		{args: args{cards: []int{1, 0, 5, 3}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCardPickup(tt.args.cards); got != tt.want {
				t.Errorf("minimumCardPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
