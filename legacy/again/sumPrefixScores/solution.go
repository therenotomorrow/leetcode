package sumPrefixScores

type TrieNode struct {
	cnt  int
	next []*TrieNode
}

func sumPrefixScores(words []string) []int {
	root := new(TrieNode)
	root.next = make([]*TrieNode, 26)

	insert := func(word string) {
		node := root
		for _, c := range word {
			if node.next[c-'a'] == nil {
				newNode := new(TrieNode)
				newNode.next = make([]*TrieNode, 26)
				node.next[c-'a'] = newNode
			}

			node.next[c-'a'].cnt++
			node = node.next[c-'a']
		}
	}

	count := func(word string) int {
		ans := 0
		node := root

		for _, c := range word {
			ans += node.next[c-'a'].cnt
			node = node.next[c-'a']
		}

		return ans
	}

	N := len(words)

	for i := range N {
		insert(words[i])
	}

	scores := make([]int, N)

	for i := range N {
		scores[i] = count(words[i])
	}

	return scores
}
