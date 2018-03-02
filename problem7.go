package projecteulergo

// Problem7 - 10001st prime
func Problem7(n int) int {
	var primes []int
	i := 0
	for j := 2; i < n; j++ {
		if isPrime(primes, j) {
			primes = append(primes, j)
			i++
		}
	}
	return primes[n-1]
}
