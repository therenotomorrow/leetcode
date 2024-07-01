package golang

import "testing"

func TestCountStudents(t *testing.T) {
	t.Parallel()

	type args struct {
		students   []int
		sandwiches []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{students: []int{1, 1, 0, 0}, sandwiches: []int{0, 1, 0, 1}}, want: 0},
		{name: "smoke 2", args: args{students: []int{1, 1, 1, 0, 0, 1}, sandwiches: []int{1, 0, 0, 0, 1, 1}}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countStudents(test.args.students, test.args.sandwiches); got != test.want {
				t.Errorf("countStudents() = %v, want = %v", got, test.want)
			}
		})
	}
}
