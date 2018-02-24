package projecteulergo

func gcd(x, y int) int {
	if x == y {
		return x
	} else if x == 1 {
		return 1
	} else if y == 1 {
		return 1
	}

	minxy := min(x, y)
	maxxy := max(x, y)
	rem := maxxy % minxy

	if 0 == rem {
		return minxy
	}

	return gcd(minxy, rem)
}

func lcm(x, y int) int {
	return (x * y) / gcd(x, y)
}

// Problem5 - Smallest multiple
func Problem5(x int) int {
	r := 1
	for i := 1; i <= x; i++ {
		r = lcm(r, i)
	}
	return r
}
