package golang

import "testing"

func TestReductionOperations(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 1, 3}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1}}, want: 0},
		{name: "smoke 3", args: args{nums: []int{1, 1, 2, 2, 3}}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reductionOperations(tt.args.nums); got != tt.want {
				t.Errorf("reductionOperations() = %v, want = %v", got, tt.want)
			}
		})
	}
}
