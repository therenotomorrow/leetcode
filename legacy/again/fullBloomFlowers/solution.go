package fullBloomFlowers

import (
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	orders := make([]int, 0, 2*len(flowers))
	canSee := make(map[int]int)

	for _, flower := range flowers {
		startToSee := flower[0]
		endToSee := flower[1] + 1

		orders = append(orders, startToSee, endToSee)

		canSee[startToSee]++
		canSee[endToSee]--
	}

	sort.Ints(orders)

	timeline := make([]int, 0)
	totalSee := make([]int, 0)
	totalNow := 0

	// add zero because of people at position 1
	timeline = append(timeline, 0)
	totalSee = append(totalSee, 0)

	for _, time := range orders {
		count, ok := canSee[time]
		if !ok {
			continue
		}

		timeline = append(timeline, time)
		totalNow += count
		totalSee = append(totalSee, totalNow)

		delete(canSee, time)
	}

	ans := make([]int, len(people))

	for i, p := range people {
		time := binarySearch(timeline, p) - 1
		ans[i] = totalSee[time]
	}

	return ans
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		m := (l + r) / 2

		if target < nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
