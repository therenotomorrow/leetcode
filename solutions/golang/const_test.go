package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

//goland:noinspection GoBoolExpressions
func TestConst(t *testing.T) {
	// make sure about values
	if want := 1000000007; golang.MOD != want {
		t.Errorf("MOD() = %v, want = %v", golang.MOD, want)
	}

	if want := -214748364; golang.MinInt32Overflow != want {
		t.Errorf("MinInt32Overflow() = %v, want = %v", golang.MinInt32Overflow, want)
	}

	if want := 214748364; golang.MaxInt32Overflow != want {
		t.Errorf("MaxInt32Overflow() = %v, want = %v", golang.MaxInt32Overflow, want)
	}
}
