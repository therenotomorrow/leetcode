package golang

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	t.Parallel()

	type args struct {
		names   []string
		heights []int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{
				names:   []string{"Mary", "John", "Emma"},
				heights: []int{180, 165, 170},
			},
			want: []string{"Mary", "Emma", "John"},
		},
		{
			name: "smoke 2",
			args: args{
				names:   []string{"Alice", "Bob", "Bob"},
				heights: []int{155, 185, 150},
			},
			want: []string{"Bob", "Alice", "Bob"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sortPeople(test.args.names, test.args.heights); !reflect.DeepEqual(got, test.want) {
				t.Errorf("sortPeople() = %v, want = %v", got, test.want)
			}
		})
	}
}
