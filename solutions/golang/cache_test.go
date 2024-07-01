package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func TestCacheSaveLoad(t *testing.T) {
	t.Parallel()

	cache := golang.NewCache()

	// empty cache
	if got, ok := cache.Load(1, 2, 3); got != 0 || ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}

	cache.Save(42, 1, 2, 3)

	// regular usage
	if got, ok := cache.Load(1, 2, 3); got != 42 || !ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	// key doesn't exist
	if got, ok := cache.Load(4, 5, 6); got != 0 || ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}

	cache.Save(100, 1, 2, 3)

	// override key
	if got, ok := cache.Load(1, 2, 3); got != 100 || !ok {
		t.Errorf("Load() = (%v, %v), want = (%v, %v)", got, ok, 100, true)
	}
}
