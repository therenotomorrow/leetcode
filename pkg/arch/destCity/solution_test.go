package destCity

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}}, want: "Sao Paulo"},
		{args: args{paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}}, want: "A"},
		{args: args{paths: [][]string{{"A", "Z"}}}, want: "Z"},
		{args: args{paths: [][]string{{"jMgaf WaWA", "iinynVdmBz"}, {" QCrEFBcAw", "wRPRHznLWS"}, {"iinynVdmBz", "OoLjlLFzjz"}, {"OoLjlLFzjz", " QCrEFBcAw"}, {"IhxjNbDeXk", "jMgaf WaWA"}, {"jmuAYy vgz", "IhxjNbDeXk"}}}, want: "wRPRHznLWS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
