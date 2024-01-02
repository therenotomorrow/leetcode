package solutions

import (
	"testing"
)

func TestFoodRatingsSmoke1(t *testing.T) {
	obj := FoodRatingsConstructor(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7},
	)
	want := ""

	want = "kimchi"

	t.Run("", func(t *testing.T) {
		if got := obj.HighestRated("korean"); got != want {
			t.Errorf(" obj.HighestRated() = %v, want = %v", got, want)
		}
	})

	want = "ramen"

	t.Run("", func(t *testing.T) {
		if got := obj.HighestRated("japanese"); got != want {
			t.Errorf(" obj.HighestRated() = %v, want = %v", got, want)
		}
	})

	obj.ChangeRating("sushi", 16)

	want = "sushi"

	t.Run("", func(t *testing.T) {
		if got := obj.HighestRated("japanese"); got != want {
			t.Errorf(" obj.HighestRated() = %v, want = %v", got, want)
		}
	})

	obj.ChangeRating("ramen", 16)

	want = "ramen"

	t.Run("", func(t *testing.T) {
		if got := obj.HighestRated("japanese"); got != want {
			t.Errorf(" obj.HighestRated() = %v, want = %v", got, want)
		}
	})
}
