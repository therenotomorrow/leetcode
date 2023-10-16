package structs_test

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"testing"
)

//goland:noinspection GoBoolExpressions
func TestConst(t *testing.T) {
	// make sure about values

	if want := 1000000007; structs.MOD != want {
		t.Errorf("MOD() = %v, want = %v", structs.MOD, want)
	}

	if want := -214748364; structs.MinInt32Overflow != want {
		t.Errorf("MinInt32Overflow() = %v, want = %v", structs.MinInt32Overflow, want)
	}

	if want := 214748364; structs.MaxInt32Overflow != want {
		t.Errorf("MaxInt32Overflow() = %v, want = %v", structs.MaxInt32Overflow, want)
	}
}
