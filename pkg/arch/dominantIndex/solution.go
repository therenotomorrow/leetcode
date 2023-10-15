package dominantIndex

func dominantIndex(nums []int) int {
	var ans, last, curr int

	for i, num := range nums {
		if num >= curr {
			last = curr
			curr = num
			ans = i
		} else if num >= last {
			last = num
		}
	}

	if curr >= 2*last {
		return ans
	}

	return -1
}
