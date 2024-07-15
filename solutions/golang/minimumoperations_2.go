package golang

func minimumOperations2(nums []int) int {
	cnt := 0

	for _, num := range nums {
		if num%3 != 0 {
			cnt++
		}
	}

	return cnt
}
