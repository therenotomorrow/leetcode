package golang

import "sort"

func smallestDistancePair(nums []int, k int) int {
	const half = 2

	sort.Ints(nums)

	minDist := 0
	maxDist := nums[len(nums)-1] - nums[0]

	for minDist < maxDist {
		midDist := (maxDist + minDist) / half
		cntDist := 0
		leftIdx := 0

		for right, num := range nums {
			for num-nums[leftIdx] > midDist {
				leftIdx++
			}

			cntDist += right - leftIdx
		}

		if cntDist < k {
			minDist = midDist + 1
		} else {
			maxDist = midDist
		}
	}

	return minDist
}
