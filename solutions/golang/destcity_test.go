package golang

import "testing"

func TestDestCity(t *testing.T) {
	t.Parallel()

	type args struct {
		paths [][]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{
				paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			},
			want: "Sao Paulo",
		},
		{
			name: "smoke 2",
			args: args{
				paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			},
			want: "A",
		},
		{
			name: "smoke 3",
			args: args{
				paths: [][]string{{"A", "Z"}},
			},
			want: "Z",
		},
		{
			name: "test 37: wrong answer",
			args: args{
				paths: [][]string{
					{"jMgaf WaWA", "iinynVdmBz"},
					{" QCrEFBcAw", "wRPRHznLWS"},
					{"iinynVdmBz", "OoLjlLFzjz"},
					{"OoLjlLFzjz", " QCrEFBcAw"},
					{"IhxjNbDeXk", "jMgaf WaWA"},
					{"jmuAYy vgz", "IhxjNbDeXk"},
				},
			},
			want: "wRPRHznLWS",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := destCity(test.args.paths); got != test.want {
				t.Errorf("destCity() = %v, want = %v", got, test.want)
			}
		})
	}
}
