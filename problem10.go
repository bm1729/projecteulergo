package projecteulergo

// Problem10 - Summation of primes
func Problem10(max int) int {
	var primes []int
	sum := 0
	for i := 2; i < max; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
			sum += i
		}
	}
	return sum
}
