package golang

import "math"

func minSwaps1(nums []int) int {
	cnt := math.MaxInt
	currSum := nums[0]
	windowSize := Sum(nums...)

	for left, right := 0, 0; left < len(nums); left++ {
		if left != 0 {
			currSum -= nums[left-1]
		}

		for right-left+1 < windowSize {
			right++
			currSum += nums[right%len(nums)]
		}

		cnt = Min(cnt, windowSize-currSum)
	}

	return cnt
}
