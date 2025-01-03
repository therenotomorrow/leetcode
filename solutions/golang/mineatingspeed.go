package golang

func minEatingSpeed(piles []int, wantH int) int {
	const half = 2

	left := 1
	right := Max(piles...)

	for left < right {
		needH := 0
		speed := left + (right-left)/half

		for _, pile := range piles {
			needH += (pile + speed - 1) / speed
		}

		if needH > wantH {
			left = speed + 1
		} else {
			right = speed
		}
	}

	return left
}
