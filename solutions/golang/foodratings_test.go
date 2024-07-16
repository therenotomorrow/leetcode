package golang

import (
	"testing"
)

func TestFoodRatingsSmoke1(t *testing.T) {
	t.Parallel()

	obj := FoodRatingsConstructor(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7},
	)

	want := "kimchi"
	if got := obj.HighestRated("korean"); got != want {
		t.Errorf("obj.HighestRated() = %v, want = %v", got, want)
	}

	want = "ramen"
	if got := obj.HighestRated("japanese"); got != want {
		t.Errorf("obj.HighestRated() = %v, want = %v", got, want)
	}

	obj.ChangeRating("sushi", 16)

	want = "sushi"
	if got := obj.HighestRated("japanese"); got != want {
		t.Errorf("obj.HighestRated() = %v, want = %v", got, want)
	}

	obj.ChangeRating("ramen", 16)

	want = "ramen"
	if got := obj.HighestRated("japanese"); got != want {
		t.Errorf("obj.HighestRated() = %v, want = %v", got, want)
	}
}
