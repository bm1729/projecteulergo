package projecteulergo

// Problem5 - Smallest multiple
func Problem5(x int) int {
	r := 1
	for i := 1; i <= x; i++ {
		r = lcm(r, i)
	}
	return r
}
