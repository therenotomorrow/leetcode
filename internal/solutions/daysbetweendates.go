package solutions

import "time"

const (
	format     = "2006-01-02"
	hoursInDay = 24
)

func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse(format, date1)
	t2, _ := time.Parse(format, date2)

	return int(t2.Sub(t1).Abs().Hours()) / hoursInDay
}
