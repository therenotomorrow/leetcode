package getAverages

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}

	div := 2*k + 1
	sum := 0
	for i := k; i < len(nums)-k; i++ {
		if sum == 0 {
			for _, num := range nums[i-k : i+k+1] {
				sum += num
			}
		} else {
			sum += nums[i+k] - nums[i-k-1]
		}

		res[i] = sum / div
	}

	return res
}
