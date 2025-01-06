package golang

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-")

	year, _ := strconv.ParseInt(parts[0], 10, 64)
	month, _ := strconv.ParseInt(parts[1], 10, 64)
	day, _ := strconv.ParseInt(parts[2], 10, 64)

	parts[0] = strconv.FormatInt(year, 2)
	parts[1] = strconv.FormatInt(month, 2)
	parts[2] = strconv.FormatInt(day, 2)

	return strings.Join(parts, "-")
}
