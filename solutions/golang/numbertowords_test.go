package golang

import "testing"

func TestNumberToWords(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{num: 123},
			want: "One Hundred Twenty Three",
		},
		{
			name: "smoke 2",
			args: args{num: 12345},
			want: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name: "smoke 3",
			args: args{num: 1234567},
			want: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name: "test 522: wrong answer",
			args: args{num: 10},
			want: "Ten",
		},
		{
			name: "test 532: wrong answer",
			args: args{num: 11},
			want: "Eleven",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberToWords(test.args.num); got != test.want {
				t.Errorf("numberToWords() = %v, want = %v", got, test.want)
			}
		})
	}
}
