package simplifyPath

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{path: "/home/"}, want: "/home"},
		{args: args{path: "/../"}, want: "/"},
		{args: args{path: "/home//foo/"}, want: "/home/foo"},
		{args: args{path: "/a/./b/../../c/"}, want: "/c"},
		{args: args{path: "/a//b////c/d/./.."}, want: "/a/b/c"},
		{args: args{path: "/..."}, want: "/..."},
		{args: args{path: "/."}, want: "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
