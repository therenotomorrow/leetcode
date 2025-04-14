package golang

func countGoodTriplets(arr []int, aParam int, bParam int, cParam int) int {
	cnt := 0

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if Abs(arr[i]-arr[j]) <= aParam && Abs(arr[j]-arr[k]) <= bParam && Abs(arr[i]-arr[k]) <= cParam {
					cnt++
				}
			}
		}
	}

	return cnt
}
