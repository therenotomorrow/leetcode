package golang

import "testing"

func TestSolution(t *testing.T) {
	t.Parallel()

	type args struct {
		file []byte
		n    int
	}

	tests := []struct {
		name    string
		args    args
		wantLen int
		wantStr string
	}{
		{name: "smoke 1", args: args{file: []byte("abc"), n: 4}, wantLen: 3, wantStr: "abc"},
		{name: "smoke 2", args: args{file: []byte("abcde"), n: 5}, wantLen: 5, wantStr: "abcde"},
		{name: "smoke 3", args: args{file: []byte("abcdABCD1234"), n: 12}, wantLen: 12, wantStr: "abcdABCD1234"},
		{name: "test 27: wrong answer", args: args{file: []byte("leetcode"), n: 5}, wantLen: 5, wantStr: "leetc"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			pnt := 0

			readFunc := solution(func(bytes []byte) int {
				var n int

				for i := 0; i < 4 && pnt < len(test.args.file); i++ {
					bytes[i] = test.args.file[pnt]
					pnt++
					n++
				}

				return n
			})

			buf := make([]byte, test.args.n)
			if len(test.args.file) < test.args.n {
				buf = make([]byte, len(test.args.file))
			}

			got := readFunc(buf, test.args.n)

			if string(buf) != test.wantStr {
				t.Errorf("solution() = %v, want = %v", string(buf), test.wantStr)
			}

			if got != test.wantLen {
				t.Errorf("solution() = %v, want = %v", got, test.wantLen)
			}
		})
	}
}
