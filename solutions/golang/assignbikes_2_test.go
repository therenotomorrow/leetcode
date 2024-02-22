package golang

import "testing"

func TestAssignBikes2(t *testing.T) {
	type args struct {
		workers [][]int
		bikes   [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				workers: [][]int{{0, 0}, {2, 1}},
				bikes:   [][]int{{1, 2}, {3, 3}},
			},
			want: 6,
		},
		{
			name: "smoke 2",
			args: args{
				workers: [][]int{{0, 0}, {1, 1}, {2, 0}},
				bikes:   [][]int{{1, 0}, {2, 2}, {2, 1}},
			}, want: 4,
		},
		{
			name: "smoke 3",
			args: args{
				workers: [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
				bikes:   [][]int{{0, 999}, {1, 999}, {2, 999}, {3, 999}, {4, 999}},
			},
			want: 4995,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignBikes2(tt.args.workers, tt.args.bikes); got != tt.want {
				t.Errorf("assignBikes2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
