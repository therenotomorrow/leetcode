package golang

func rearrangeArray(nums []int) []int {
	var (
		i, j, k int
		arr     = make([]int, len(nums))
	)

	for k < len(arr) {
		// search first positive
		for nums[i] < 0 && i < len(nums) {
			i++
		}

		// search first negative
		for nums[j] > 0 && j < len(nums) {
			j++
		}

		arr[k] = nums[i]
		k++

		arr[k] = nums[j]
		k++

		i++
		j++
	}

	return arr
}
