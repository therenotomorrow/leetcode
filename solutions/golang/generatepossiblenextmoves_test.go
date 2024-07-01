package golang

import (
	"reflect"
	"testing"
)

func TestGeneratePossibleNextMoves(t *testing.T) {
	t.Parallel()

	type args struct {
		currentState string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{currentState: "++++"}, want: []string{"--++", "+--+", "++--"}},
		{name: "smoke 2", args: args{currentState: "+"}, want: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := generatePossibleNextMoves(test.args.currentState); !reflect.DeepEqual(got, test.want) {
				t.Errorf("generatePossibleNextMoves() = %v, want = %v", got, test.want)
			}
		})
	}
}
