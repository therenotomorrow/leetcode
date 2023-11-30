package minimumOneBitOperations

import "testing"

func Test_minimumOneBitOperations(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 3}, want: 2},
		{args: args{n: 6}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOneBitOperations(tt.args.n); got != tt.want {
				t.Errorf("minimumOneBitOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
