package golang

import "sort"

func triangleType(nums []int) string {
	const (
		equilateral = "equilateral"
		isosceles   = "isosceles"
		scalene     = "scalene"
		none        = "none"
	)

	if nums[0]+nums[1] <= nums[2] || nums[0]+nums[2] <= nums[1] || nums[1]+nums[2] <= nums[0] {
		return none
	}

	sort.Ints(nums)

	switch {
	case nums[0] == nums[2]:
		return equilateral
	case nums[0] == nums[1] || nums[1] == nums[2]:
		return isosceles
	default:
		return scalene
	}
}
