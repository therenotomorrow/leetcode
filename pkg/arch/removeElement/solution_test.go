package removeElement

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantArr []int
	}{
		{args: args{nums: []int{3, 2, 2, 3}, val: 3}, want: 2, wantArr: []int{2, 2}},
		{args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, want: 5, wantArr: []int{0, 1, 3, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.nums[:tt.want], tt.wantArr) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums[:tt.want], tt.wantArr)
			}
		})
	}
}
