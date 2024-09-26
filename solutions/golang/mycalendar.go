package golang

type MyCalendar struct {
	events [][]int
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{events: make([][]int, 0)}
}

func (m *MyCalendar) Book(start int, end int) bool {
	for _, event := range m.events {
		if event[0] < end && start < event[1] {
			return false
		}
	}

	m.events = append(m.events, []int{start, end})

	return true
}
