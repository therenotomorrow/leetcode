package solutions

import (
	"reflect"
	"testing"
)

func TestGeneratePossibleNextMoves(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePossibleNextMoves(tt.args.currentState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePossibleNextMoves() = %v, want = %v", got, tt.want)
			}
		})
	}
}
