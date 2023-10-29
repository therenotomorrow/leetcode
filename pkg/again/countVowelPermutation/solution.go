package countVowelPermutation

const MOD int = 1e9 + 7

func countVowelPermutation(n int) int {
	// order: 'a', 'e', 'i', 'o', 'u'
	endsWithN := [...]int{1, 1, 1, 1, 1}

	for i := 1; i < n; i++ {
		endsWithN = [...]int{
			Sum(endsWithN[1], endsWithN[2], endsWithN[4]) % MOD,
			Sum(endsWithN[0], endsWithN[2]) % MOD,
			Sum(endsWithN[1], endsWithN[3]) % MOD,
			Sum(endsWithN[2]) % MOD,
			Sum(endsWithN[2], endsWithN[3]) % MOD,
		}
	}

	return Sum(endsWithN[0], endsWithN[1:]...) % MOD
}

func Sum(num int, nums ...int) int {
	s := num
	for _, n := range nums {
		s += n
	}
	return s
}
