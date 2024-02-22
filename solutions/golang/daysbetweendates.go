package golang

import "time"

func daysBetweenDates(date1 string, date2 string) int {
	const (
		format     = "2006-01-02"
		hoursInDay = 24
	)

	t1, _ := time.Parse(format, date1)
	t2, _ := time.Parse(format, date2)

	return int(t2.Sub(t1).Abs().Hours()) / hoursInDay
}
