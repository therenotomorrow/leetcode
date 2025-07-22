package golang

func maximumUniqueSubarray(nums []int) int {
	ans := 0
	sum := 0
	left := 0
	cnt := make(map[int]int)

	for right, num := range nums {
		if last, ok := cnt[num]; ok {
			for ; left <= last; left++ {
				sum -= nums[left]

				delete(cnt, nums[left])
			}
		}

		sum += num
		if sum > ans {
			ans = sum
		}

		cnt[num] = right
	}

	return ans
}
