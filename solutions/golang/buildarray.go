package golang

func buildArray(target []int, _ int) []string {
	ops := make([]string, 0)
	stream := 1

	for _, num := range target {
		for ; stream != num; stream++ {
			ops = append(ops, "Push", "Pop")
		}

		ops = append(ops, "Push")
		stream++
	}

	return ops
}
