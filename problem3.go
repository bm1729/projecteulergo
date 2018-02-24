package projecteulergo

// Problem3 - Largest prime factor
func Problem3(x int) int {
	d := 2
	for d < x {
		if x%d == 0 {
			x = x / d
		} else {
			d++
		}
	}
	return x
}
