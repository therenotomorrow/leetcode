package summaryRanges

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0)

	i := nums[0]
	l, r := nums[0], i

	for _, num := range nums {
		if num != i {
			if l == r {
				res = append(res, strconv.Itoa(l))
			} else {
				res = append(res, strconv.Itoa(l)+"->"+strconv.Itoa(r))
			}

			l = num
			i = num
		}

		r = num
		i++
	}

	if l == r {
		res = append(res, strconv.Itoa(l))
	} else {
		res = append(res, strconv.Itoa(l)+"->"+strconv.Itoa(r))
	}

	return res
}
