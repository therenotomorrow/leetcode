package golang

func findComplement(num int) int {
	powered := 1

	for powered <= num {
		powered *= 2
	}

	return num ^ (powered - 1)
}
