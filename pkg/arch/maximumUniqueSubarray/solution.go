package maximumUniqueSubarray

func maximumUniqueSubarray(nums []int) int {
	var (
		ans, sum int
		counter  = make(map[int]int)
	)

	for l, r := 0, 0; r < len(nums); r++ {
		last, ok := counter[nums[r]]
		if ok {
			for ; l <= last; l++ {
				sum -= nums[l]

				delete(counter, nums[l])
			}
		}

		sum += nums[r]
		if sum > ans {
			ans = sum
		}

		counter[nums[r]] = r
	}

	return ans
}
