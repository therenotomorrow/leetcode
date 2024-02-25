package canTraverseAllPairs

type DSU struct {
	dsu  []int
	size []int
}

func (d *DSU) find(x int) int {
	if d.dsu[x] == x {
		return x
	}

	d.dsu[x] = d.find(d.dsu[x])
	return d.dsu[x]
}

func (d *DSU) merge(x int, y int) {
	fx := d.find(x)
	fy := d.find(y)
	if fx == fy {
		return
	}

	if d.size[fx] > d.size[fy] {
		fx, fy = fy, fx
	}

	d.dsu[fx] = fy
	d.size[fy] += d.size[fx]
}

func NewDSU(n int) *DSU {
	dsu := make([]int, n+1)
	size := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dsu[i] = i
		size[i] = 1
	}

	return &DSU{dsu: dsu, size: size}
}

func canTraverseAllPairs(nums []int) bool {
	MAX := 100000
	n := len(nums)

	has := make([]bool, MAX+1)

	for _, x := range nums {
		has[x] = true
	}

	// edge cases
	if n == 1 {
		return true
	}
	if has[1] {
		return false
	}

	// the general solution
	sieve := make([]int, MAX+1)
	for d := 2; d <= MAX; d++ {
		if sieve[d] == 0 {
			for v := d; v <= MAX; v += d {
				sieve[v] = d
			}
		}
	}

	union := NewDSU(2*MAX + 1)

	for _, x := range nums {
		val := x
		for val > 1 {
			prime := sieve[val]
			root := prime + MAX
			if union.find(root) != union.find(x) {
				union.merge(root, x)
			}
			for val%prime == 0 {
				val /= prime
			}
		}
	}

	cnt := 0
	for i := 2; i <= MAX; i++ {
		if has[i] && union.find(i) == i {
			cnt++
		}
	}

	return cnt == 1
}
