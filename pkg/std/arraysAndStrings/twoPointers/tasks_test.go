package twoPointers

import (
	"reflect"
	"testing"
)

func Test_checkForTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nums: []int{1, 2, 4, 6, 8, 9, 14, 15}, target: 13}, want: true},
		{args: args{nums: []int{1, 2, 4, 6}, target: 13}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkForTarget(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("checkForTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIfPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "abcdcba"}, want: true},
		{args: args{s: "racecar"}, want: true},
		{args: args{s: "haha"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPalindrome(tt.args.s); got != tt.want {
				t.Errorf("checkIfPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combine(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{arr1: []int{1, 4, 7, 20}, arr2: []int{3, 5, 6}}, want: []int{1, 3, 4, 5, 6, 7, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
