package RecentCounter

type RecentCounter struct {
	reqs []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	l := t - 3000
	r := t

	rc.reqs = append(rc.reqs, t)

	j := -1
	for i, req := range rc.reqs {
		if l <= req && req <= r {
			break
		}

		j = i
	}

	if j != -1 {
		rc.reqs = rc.reqs[j+1:]
	}

	return len(rc.reqs)
}
