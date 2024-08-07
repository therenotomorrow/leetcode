package golang

import (
	"slices"
	"strings"
)

func groupAdder(target []string, group []string, i int) []string {
	if group[i] == "" {
		return target
	}

	return append(target, group[i])
}

func numberToWords(num int) string {
	const (
		maxGroupSize = 4
		tens         = 10
		twenties     = 20
		hundreds     = 100
		groupDigits  = 1000
	)

	var (
		groupOfSimple = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
		groupOfTeen   = []string{
			"",
			"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
		}
		groupOfTy     = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		groupOfDigits = []string{"", "Thousand", "Million", "Billion"}
		digitGroup    = -1
		groups        = make([]string, 0)
	)

	// simple case
	if num == 0 {
		return "Zero"
	}

	for ; num > 0; num /= groupDigits {
		digitGroup++
		part := num % groupDigits

		// the whole group of zeros here
		if part == 0 {
			continue
		}

		group := make([]string, 0, maxGroupSize)

		if part >= hundreds {
			group = groupAdder(group, groupOfSimple, part/hundreds)
			group = append(group, "Hundred")
			part %= hundreds
		}

		if part >= twenties {
			group = groupAdder(group, groupOfTy, part/tens)
			part %= tens
		}

		if part > tens {
			group = groupAdder(group, groupOfTeen, part%tens)
		} else {
			group = groupAdder(group, groupOfSimple, part)
		}

		groups = groupAdder(groups, groupOfDigits, digitGroup)
		groups = append(groups, strings.Join(group, " "))
	}

	slices.Reverse(groups)

	return strings.Join(groups, " ")
}
