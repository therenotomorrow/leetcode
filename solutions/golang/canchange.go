package golang

func canChange(start string, target string) bool { //nolint:cyclop
	i := 0
	j := 0

	for i < len(start) || j < len(target) {
		for i < len(start) && start[i] == '_' {
			i++
		}

		for j < len(target) && target[j] == '_' {
			j++
		}

		if i == len(start) || j == len(target) {
			return i == len(start) && j == len(target)
		}

		if start[i] != target[j] || start[i] == 'L' && i < j || start[i] == 'R' && i > j {
			return false
		}

		i++
		j++
	}

	return true
}
