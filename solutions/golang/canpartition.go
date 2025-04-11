package golang

func canPartition(nums []int) bool {
	sum := Sum(nums...)

	if sum%Half != 0 {
		return false
	}

	var (
		cache   = make([]map[int]bool, len(nums)+1)
		dynamic func(curr int, sum int) bool
	)

	for key := range cache {
		cache[key] = make(map[int]bool)
	}

	dynamic = func(curr int, sum int) bool {
		if sum < 0 || curr >= len(nums) {
			return false
		}

		if sum == 0 {
			return true
		}

		if val, ok := cache[curr][sum]; ok {
			return val
		}

		res := dynamic(curr+1, sum-nums[curr]) || dynamic(curr+1, sum)

		cache[curr][sum] = res

		return res
	}

	return dynamic(0, sum/Half)
}
