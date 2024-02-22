package solutions

import "testing"

func TestFindJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 2, trust: [][]int{{1, 2}}}, want: 2},
		{name: "smoke 2", args: args{n: 3, trust: [][]int{{1, 3}, {2, 3}}}, want: 3},
		{name: "smoke 3", args: args{n: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}}, want: -1},
		{name: "test 4: wrong answer", args: args{n: 4, trust: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}}, want: 3},
		{name: "test 87: wrong answer", args: args{n: 3, trust: [][]int{{1, 2}, {2, 3}}}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want = %v", got, tt.want)
			}
		})
	}
}
