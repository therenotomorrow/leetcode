package golang

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	versions1 := strings.Split(version1, ".")
	versions2 := strings.Split(version2, ".")
	limit := Max(len(versions1), len(versions2))

	for part := 0; part < limit; part++ {
		var subVer1, subVer2 int

		if part < len(versions1) {
			subVer1, _ = strconv.Atoi(versions1[part])
		}

		if part < len(versions2) {
			subVer2, _ = strconv.Atoi(versions2[part])
		}

		if subVer1 != subVer2 {
			if subVer1 > subVer2 {
				return 1
			}

			return -1
		}
	}

	return 0
}
