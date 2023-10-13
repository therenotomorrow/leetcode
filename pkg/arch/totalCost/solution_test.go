package totalCost

import "testing"

func Test_totalCost(t *testing.T) {
	type args struct {
		costs      []int
		k          int
		candidates int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{args: args{costs: []int{17, 12, 10, 2, 7, 2, 11, 20, 8}, k: 3, candidates: 4}, want: 11},
		{args: args{costs: []int{1, 2, 4, 1}, k: 3, candidates: 3}, want: 4},
		{args: args{costs: []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, k: 11, candidates: 2}, want: 423},
		{args: args{costs: []int{28, 35, 21, 13, 21, 72, 35, 52, 74, 92, 25, 65, 77, 1, 73, 32, 43, 68, 8, 100, 84, 80, 14, 88, 42, 53, 98, 69, 64, 40, 60, 23, 99, 83, 5, 21, 76, 34}, k: 32, candidates: 12}, want: 1407},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalCost(tt.args.costs, tt.args.k, tt.args.candidates); got != tt.want {
				t.Errorf("totalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
