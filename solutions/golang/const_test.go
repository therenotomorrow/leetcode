package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

//goland:noinspection GoBoolExpressions
func TestConst(t *testing.T) { //nolint:cyclop
	t.Parallel()

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

	if want := "smoke 1"; golang.Smoke1 != want {
		t.Errorf("Smoke1() = %v, want = %v", golang.Smoke1, want)
	}

	if want := "smoke 2"; golang.Smoke2 != want {
		t.Errorf("Smoke2() = %v, want = %v", golang.Smoke2, want)
	}

	if want := "smoke 3"; golang.Smoke3 != want {
		t.Errorf("Smoke3() = %v, want = %v", golang.Smoke3, want)
	}

	if want := 10; golang.Digits != want {
		t.Errorf("Digits() = %v, want = %v", golang.Digits, want)
	}

	if want := 26; golang.Alphabet != want {
		t.Errorf("Alphabet() = %v, want = %v", golang.Alphabet, want)
	}

	if want := 255; golang.Byte != want {
		t.Errorf("Byte() = %v, want = %v", golang.Byte, want)
	}

	if want := 7; golang.Base7 != want {
		t.Errorf("Base7() = %v, want = %v", golang.Base7, want)
	}

	if want := 2; golang.Base2 != want {
		t.Errorf("Base2() = %v, want = %v", golang.Base2, want)
	}

	if want := 2; golang.Half != want {
		t.Errorf("Half() = %v, want = %v", golang.Half, want)
	}

	if want := 2; golang.Double != want {
		t.Errorf("Double() = %v, want = %v", golang.Double, want)
	}
}
