package golang

import "testing"

func TestIsStrobogrammatic(t *testing.T) {
	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{num: "69"}, want: true},
		{name: "smoke 2", args: args{num: "88"}, want: true},
		{name: "smoke 3", args: args{num: "962"}, want: false},
		{name: "test 30: wrong answer", args: args{num: "2"}, want: false},
		{name: "test 38: wrong answer", args: args{num: "3"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStrobogrammatic(tt.args.num); got != tt.want {
				t.Errorf("isStrobogrammatic() = %v, want = %v", got, tt.want)
			}
		})
	}
}
