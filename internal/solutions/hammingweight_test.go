package solutions

import "testing"

func TestHammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{num: 00000000000000000000000000001011}, want: 3},
		{name: "smoke 2", args: args{num: 00000000000000000000000010000000}, want: 1},
		{name: "smoke 3", args: args{num: 00000000000000000000010000100101}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want = %v", got, tt.want)
			}
		})
	}
}
