package solutions

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	curr := 0

	for right := range nums {
		target := nums[right]
		curr += target

		if (right-left+1)*target-curr > k {
			curr -= nums[left]
			left++
		}
	}

	return len(nums) - left
}
