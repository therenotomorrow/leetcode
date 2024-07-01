package golang

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	const diff = 2

	if len(nums) == 0 {
		return [][]int{{lower, upper}}
	}

	ranges := make([][]int, 0)
	if lower < nums[0] {
		ranges = append(ranges, []int{lower, nums[0] - 1})
	}

	for i := range len(nums) - 1 {
		if nums[i+1]-nums[i] < diff {
			continue
		}

		ranges = append(ranges, []int{nums[i] + 1, nums[i+1] - 1})
	}

	if upper > nums[len(nums)-1] {
		ranges = append(ranges, []int{nums[len(nums)-1] + 1, upper})
	}

	return ranges
}
