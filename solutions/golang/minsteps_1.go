package golang

func minSteps1(str string, tpl string) int {
	cnt := make(map[byte]int)

	// les(str) == len(tpl) by description
	for i := range len(str) {
		cnt[str[i]]++
		cnt[tpl[i]]--
	}

	steps := 0

	for _, times := range cnt {
		if times < 0 {
			steps += -times
		}
	}

	return steps
}
