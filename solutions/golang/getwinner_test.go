package golang

import "testing"

func TestGetWinner(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{2, 1, 3, 5, 4, 6, 7}, k: 2}, want: 5},
		{name: "smoke 2", args: args{arr: []int{3, 2, 1}, k: 10}, want: 3},
		{name: "test 3: time limit", args: args{arr: []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, k: 1000000000}, want: 99},
		{name: "test 139: wrong answer", args: args{arr: []int{1, 25, 35, 42, 68, 70}, k: 2}, want: 70},
		{name: "test 160: wrong answer", args: args{arr: []int{1, 25, 68, 35, 42, 70}, k: 2}, want: 68},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinner(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getWinner() = %v, want = %v", got, tt.want)
			}
		})
	}
}
