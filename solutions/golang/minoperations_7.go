package golang

func minOperations7(nums []int) int {
	cnt := 0

	for i := 2; i < len(nums); i++ {
		if nums[i-2] == 0 {
			cnt++

			nums[i-2] ^= 1
			nums[i-1] ^= 1
			nums[i] ^= 1
		}
	}

	if Sum(nums...) == len(nums) {
		return cnt
	}

	return -1
}
