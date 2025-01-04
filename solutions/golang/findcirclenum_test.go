package golang

import "testing"

func TestFindCircleNum(t *testing.T) {
	t.Parallel()

	type args struct {
		isConnected [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}}, want: 2},
		{name: "smoke 2", args: args{isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findCircleNum(test.args.isConnected); got != test.want {
				t.Errorf("findCircleNum() = %v, want = %v", got, test.want)
			}
		})
	}
}
