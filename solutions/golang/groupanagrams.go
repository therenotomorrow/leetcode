package golang

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		key := []rune(str)

		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})

		anagrams[string(key)] = append(anagrams[string(key)], str)
	}

	res := make([][]string, 0)

	for _, vals := range anagrams {
		res = append(res, vals)
	}

	return res
}
