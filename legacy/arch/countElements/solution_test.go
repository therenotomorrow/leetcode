package countElements

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{arr: []int{1, 2, 3}}, want: 2},
		{args: args{arr: []int{1, 1, 3, 3, 5, 5, 7, 7}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.arr); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
