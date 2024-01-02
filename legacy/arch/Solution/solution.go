package Solution

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (s *Solution) Reset() []int {
	return s.nums
}

func (s *Solution) Shuffle() []int {
	nums := make([]int, len(s.nums))

	copy(nums, s.nums)

	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	return nums
}
