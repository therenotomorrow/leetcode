package golang

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
