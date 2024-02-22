package golang

import "testing"

func TestSolution(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pnt := 0

			readFunc := solution(func(bytes []byte) int {
				var n int

				for i := 0; i < 4 && pnt < len(tt.args.file); i++ {
					bytes[i] = tt.args.file[pnt]
					pnt++
					n++
				}

				return n
			})

			buf := make([]byte, tt.args.n)
			if len(tt.args.file) < tt.args.n {
				buf = make([]byte, len(tt.args.file))
			}

			got := readFunc(buf, tt.args.n)

			if string(buf) != tt.wantStr {
				t.Errorf("solution() = %v, want = %v", string(buf), tt.wantStr)
			}

			if got != tt.wantLen {
				t.Errorf("solution() = %v, want = %v", got, tt.wantLen)
			}
		})
	}
}
