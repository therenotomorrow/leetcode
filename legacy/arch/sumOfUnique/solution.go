package sumOfUnique

func sumOfUnique(nums []int) int {
	arr := make([]int, 101)

	for _, num := range nums {
		arr[num]++
	}

	sum := 0
	for num, cnt := range arr {
		if cnt == 1 {
			sum += num
		}
	}

	return sum
}
