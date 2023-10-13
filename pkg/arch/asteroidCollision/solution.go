package asteroidCollision

func isCollide(a int, b int) bool {
	return a > 0 && b < 0
}

func asteroidCollision(asteroids []int) []int {
	stack := []int{asteroids[0]}

	for _, aster := range asteroids[1:] {
		if len(stack) < 1 || !isCollide(stack[len(stack)-1], aster) {
			stack = append(stack, aster)
			continue
		}

		for {
			if len(stack) == 0 {
				stack = append(stack, aster)
				break
			} else if isCollide(stack[len(stack)-1], aster) && stack[len(stack)-1] < -aster {
				stack = stack[:len(stack)-1]
			} else if isCollide(stack[len(stack)-1], aster) && stack[len(stack)-1] == -aster {
				stack = stack[:len(stack)-1]
				break
			} else if !isCollide(stack[len(stack)-1], aster) {
				stack = append(stack, aster)
				break
			} else {
				break
			}
		}
	}

	return stack
}
