package golang

func minSteps(s string, t string) int {
	cnt := make(map[byte]int)

	// les(s) == len(t) by description
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
		cnt[t[i]]--
	}

	steps := 0

	for _, times := range cnt {
		if times < 0 {
			steps += -times
		}
	}

	return steps
}
