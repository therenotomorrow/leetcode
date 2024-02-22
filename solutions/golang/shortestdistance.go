package golang

import (
	"math"
)

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	wordIdx1 := make([]int, 0)
	wordIdx2 := make([]int, 0)

	for i, word := range wordsDict {
		switch word {
		case word1:
			wordIdx1 = append(wordIdx1, i)
		case word2:
			wordIdx2 = append(wordIdx2, i)
		}
	}

	shortest := math.MaxInt

	for _, i := range wordIdx1 {
		for _, j := range wordIdx2 {
			shortest = Min(shortest, Abs(j-i))
		}
	}

	return shortest
}
