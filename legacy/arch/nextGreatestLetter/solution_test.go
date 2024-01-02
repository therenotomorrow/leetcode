package nextGreatestLetter

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// {args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'}, want: 'c'},
		//{args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'}, want: 'f'},
		{args: args{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z'}, want: 'x'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
