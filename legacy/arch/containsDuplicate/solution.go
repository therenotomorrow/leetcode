package containsDuplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		_, exist := m[num]
		if exist {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}
