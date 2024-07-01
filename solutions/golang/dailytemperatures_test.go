package golang

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	t.Parallel()

	type args struct {
		temperatures []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{name: "smoke 2", args: args{temperatures: []int{30, 40, 50, 60}}, want: []int{1, 1, 1, 0}},
		{name: "smoke 3", args: args{temperatures: []int{30, 60, 90}}, want: []int{1, 1, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := dailyTemperatures(test.args.temperatures); !reflect.DeepEqual(got, test.want) {
				t.Errorf("dailyTemperatures() = %v, want = %v", got, test.want)
			}
		})
	}
}
