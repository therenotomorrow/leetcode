package solutions

type SeatManager struct {
	seats []int
	curr  int
}

func Constructor(n int) SeatManager {
	// could be solved using heap
	return SeatManager{seats: make([]int, n+1), curr: 1}
}

func (sm *SeatManager) Reserve() int {
	curr := sm.curr

	for i, seat := range sm.seats[curr:] {
		if seat == 0 {
			curr += i
			break
		}
	}

	sm.curr++

	sm.seats[curr] = 1

	return curr
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	sm.seats[seatNumber] = 0

	if seatNumber < sm.curr {
		sm.curr = seatNumber
	}
}
