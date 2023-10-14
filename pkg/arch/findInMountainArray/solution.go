package findInMountainArray

type MountainArray struct {
	data []int
}

func (arr *MountainArray) get(index int) int {
	return arr.data[index]
}

func (arr *MountainArray) length() int {
	return len(arr.data)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peakIdx := findPeak(mountainArr)

	// other values will be low because of peak
	if mountainArr.get(peakIdx) == target {
		return peakIdx
	}

	foundL := findLTarget(mountainArr, 0, peakIdx, target)
	foundR := findRTarget(mountainArr, peakIdx, mountainArr.length(), target)

	if foundL*foundR == 1 {
		return -1
	}

	if foundL == -1 {
		return foundR
	}

	if foundR == -1 {
		return foundL
	}

	return min(foundL, foundR)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findPeak(arr *MountainArray) int {
	peak, r, m := 0, arr.length()-1, 0

	for peak < r {
		m = (peak + r) / 2

		if arr.get(m) < arr.get(m+1) {
			peak = m + 1
		} else {
			r = m
		}
	}

	return peak
}

func findLTarget(arr *MountainArray, l int, r int, target int) int {
	for l < r {
		m := (l + r) / 2

		switch {
		case target == arr.get(m):
			return m
		case target < arr.get(m):
			r = m
		default:
			l = m + 1
		}
	}

	return -1
}

func findRTarget(arr *MountainArray, l int, r int, target int) int {
	for l < r {
		m := (l + r) / 2

		switch {
		case target == arr.get(m):
			return m
		case target < arr.get(m):
			l = m + 1
		default:
			r = m
		}
	}

	return -1
}
