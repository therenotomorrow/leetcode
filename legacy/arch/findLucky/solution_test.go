package findLucky

import "testing"

func Test_findLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{arr: []int{2, 2, 3, 4}}, want: 2},
		{args: args{arr: []int{1, 2, 2, 3, 3, 3}}, want: 3},
		{args: args{arr: []int{2, 2, 2, 3, 3}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.args.arr); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
