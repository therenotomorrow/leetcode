package golang

import "fmt"

type ValidWordAbbr struct {
	uniq map[string]bool
	data map[string]struct{}
}

func ValidWordAbbrConstructor(dictionary []string) ValidWordAbbr {
	obj := ValidWordAbbr{
		uniq: make(map[string]bool),
		data: make(map[string]struct{}),
	}

	for _, word := range dictionary {
		obj.data[word] = struct{}{}
	}

	for word := range obj.data {
		key := obj.genKey(word)

		_, ok := obj.uniq[key]

		obj.uniq[key] = !ok
	}

	return obj
}

func (abbr *ValidWordAbbr) IsUnique(word string) bool {
	key := abbr.genKey(word)

	uniq, ok := abbr.uniq[key]
	_, contains := abbr.data[word]

	// if abbr not exists or not unique
	// if abbr uniq and we stored it
	return !ok && !uniq || uniq && contains
}

func (abbr *ValidWordAbbr) genKey(word string) string {
	const minSize = 2

	n := len(word)

	if n <= minSize {
		return word
	}

	return fmt.Sprintf("%c%d%c", word[0], n-minSize, word[n-1])
}
