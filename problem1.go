package projecteulergo

// Problem1 - Multiples of 3 and 5 up to max
func Problem1(max int) int {
	sum := 0
	for i := 0; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}
