package projecteulergo

// Problem2 - Even Fibonacci numbers less than max
func Problem2(max int) int {
	sum := 0

	for x, y := 1, 2; y < max; {
		if y%2 == 0 {
			sum += y
		}
		t := x
		x = y
		y = t + y
	}

	return sum
}
