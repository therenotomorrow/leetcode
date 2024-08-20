package golang

func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)

	// mark numbers as indexes that exists in array
	for _, num := range nums {
		i := Abs(num) - 1

		if nums[i] > 0 {
			nums[i] *= -1
		}
	}

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
