package golang

func canWinNim(n int) bool {
	const maxStones = 3

	return n%(maxStones+1) != 0
}
