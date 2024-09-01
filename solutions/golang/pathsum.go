package golang

func pathSum(nums []int) int {
	const double = 2

	if len(nums) == 0 {
		return 0
	}

	var (
		coordinates = make(map[int]int)
		dfs         func(curr int, sum int) int
	)

	dfs = func(coordinate int, prevSum int) int {
		level := coordinate / Digits
		position := coordinate % Digits

		// 2 * position is a known because each level doubles coordinates
		leftCoordinate := (level+1)*Digits + double*position - 1
		rightCoordinate := (level+1)*Digits + double*position
		currSum := prevSum + coordinates[coordinate]

		// we don't have these nodes - they are leafs
		_, okLeft := coordinates[leftCoordinate]
		_, okRight := coordinates[rightCoordinate]

		if !okLeft && !okRight {
			return currSum
		}

		leftSum := 0
		if okLeft {
			leftSum = dfs(leftCoordinate, currSum)
		}

		rightSum := 0
		if okRight {
			rightSum = dfs(rightCoordinate, currSum)
		}

		return leftSum + rightSum
	}

	for _, num := range nums {
		coordinate := num / Digits
		value := num % Digits
		coordinates[coordinate] = value
	}

	return dfs(nums[0]/Digits, 0)
}
