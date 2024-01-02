package maximumWealth

import "testing"

func Test_maximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{accounts: [][]int{{1, 2, 3}, {1, 2, 3}}}, want: 6},
		{args: args{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}}, want: 10},
		{args: args{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
