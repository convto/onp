package onp

type Diff struct {
	a, b []rune
	m, n int // m <= n, m is a length, n is b length
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// NewDiff creates Diff so that m <= n.
func NewDiff(a, b []rune) Diff {
	m, n := len(a), len(b)
	if m > n {
		return Diff{a: b, b: a, m: n, n: m}
	}
	return Diff{a: a, b: b, m: m, n: n}
}

func (d Diff) ONP() int {
	offset := d.m + 1
	delta := d.n - d.m
	fp := make([]int, d.m+d.n+3)
	for i := range fp {
		fp[i] = -1
	}
	for p := 0; ; p++ {
		// -p <= k <= delta - 1
		for k := -p; k <= delta-1; k++ {
			fp[k+offset] = d.snake(k, max(fp[k-1+offset]+1, fp[k+1+offset]))
		}
		// d.delta + 1 <= k <= d.delta + p
		for k := delta + p; k >= delta+1; k-- {
			fp[k+offset] = d.snake(k, max(fp[k-1+offset]+1, fp[k+1+offset]))
		}
		// d.delta == k
		fp[delta+offset] = d.snake(delta, max(fp[delta-1+offset]+1, fp[delta+1+offset]))
		if fp[delta+offset] == d.n {
			return delta + 2*p
		}
	}
}

// snake denote the y-coordinate of the furthest point on diagonal k
func (d Diff) snake(k, y int) int {
	x := y - k
	for x < d.m && y < d.n && d.a[x] == d.b[y] {
		x++
		y++
	}
	return y
}
