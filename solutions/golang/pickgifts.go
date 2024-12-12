package golang

import (
	"math"
	"sort"
)

func pickGifts(gifts []int, k int) int64 {
	slice := sort.IntSlice(gifts)

	for ; k > 0; k-- {
		sort.Stable(sort.Reverse(slice))
		slice[0] = int(math.Sqrt(float64(gifts[0])))
	}

	return int64(Sum(slice...))
}
