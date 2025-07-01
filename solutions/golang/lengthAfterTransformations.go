package golang

func lengthAfterTransformations(text string, times int) int {
	ans := 0
	cnt := make([]int, Alphabet)

	for _, r := range text {
		cnt[r-'a']++
	}

	for range times {
		tmp := make([]int, Alphabet)

		tmp[0] = cnt[Alphabet-1]         // z -> a
		tmp[1] = (tmp[0] + cnt[0]) % MOD // a -> b and z -> ab

		for i := 2; i < Alphabet; i++ {
			tmp[i] = cnt[i-1]
		}

		cnt = tmp
	}

	for _, times := range cnt {
		ans = (ans + times) % MOD
	}

	return ans
}
