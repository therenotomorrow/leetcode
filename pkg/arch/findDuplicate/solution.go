package findDuplicate

func findDuplicate(nums []int) int {
	// Floyd's cycle
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]

		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
