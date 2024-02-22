package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func TestCacheSaveLoad(t *testing.T) {
	c := golang.NewCache()

	// empty cache
	if got, ok := c.Load(1, 2, 3); got != 0 || ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}

	c.Save(42, 1, 2, 3)

	// regular usage
	if got, ok := c.Load(1, 2, 3); got != 42 || !ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	// key doesn't exist
	if got, ok := c.Load(4, 5, 6); got != 0 || ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}

	c.Save(100, 1, 2, 3)

	// override key
	if got, ok := c.Load(1, 2, 3); got != 100 || !ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 100, true)
	}
}
