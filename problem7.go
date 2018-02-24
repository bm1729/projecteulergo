package projecteulergo

import "math"

func isPrime(primes []int, x int) bool {
	for _, prime := range primes {
		if float64(prime) > math.Sqrt(float64(x))+0.0001 { // Floating point weirdness
			break
		} else if x%prime == 0 {

			// Found a prime factor
			return false
		}
	}

	// No prime factors - must be a new prime
	return true
}

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
