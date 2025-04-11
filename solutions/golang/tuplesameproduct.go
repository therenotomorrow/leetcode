package golang

func tupleSameProduct(nums []int) int {
	const combinations = 8

	var (
		ans int
		cnt = make(map[int]int)
	)

	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			cnt[nums[j]*num]++
		}
	}

	for _, pairs := range cnt {
		ans += combinations * (pairs * (pairs - 1) / Half)
	}

	return ans
}
