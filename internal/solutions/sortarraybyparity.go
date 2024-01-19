package solutions

func sortArrayByParity(nums []int) []int {
	sorted := make([]int, len(nums))
	pnt := 0

	for _, num := range nums {
		if num%2 == 0 {
			sorted[pnt] = num
			pnt++
		}
	}

	for _, num := range nums {
		if num%2 == 1 {
			sorted[pnt] = num
			pnt++
		}
	}

	return sorted
}
