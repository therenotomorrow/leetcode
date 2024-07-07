package golang

func numWaterBottles(numBottles int, numExchange int) int {
	bottles := numBottles

	for numBottles >= numExchange {
		full := numBottles / numExchange
		rest := numBottles % numExchange

		bottles += full

		numBottles = full + rest
	}

	return bottles
}
