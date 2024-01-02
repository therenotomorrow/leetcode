package peakIndexInMountainArray

func peakIndexInMountainArray(arr []int) int {
	peak, r, m := 0, len(arr)-1, 0

	for peak < r {
		m = (peak + r) / 2

		if arr[m] < arr[m+1] {
			peak = m + 1
		} else {
			r = m
		}
	}

	return peak
}
