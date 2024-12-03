package golang

import (
	"strings"
	"time"
)

func reformatDate(date string) string {
	const (
		inputLayout  = "2 Jan 2006"
		outputLayout = "2006-01-02"
	)

	parts := strings.Fields(date)

	date = parts[0][:len(parts[0])-2] + " " + parts[1] + " " + parts[2]

	parsed, _ := time.Parse(inputLayout, date)

	return parsed.Format(outputLayout)
}
