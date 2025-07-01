package golang

func sortColors(nums []int) {
	const (
		red = iota
		_   // white
		blue
	)

	var (
		left  = 0
		right = len(nums) - 1
	)

	for curr := 0; curr <= right; {
		switch num := nums[curr]; num {
		case blue:
			nums[right], nums[curr] = nums[curr], nums[right]
			right--
		case red:
			nums[left], nums[curr] = nums[curr], nums[left]
			left++

			fallthrough
		default:
			curr++
		}
	}
}
