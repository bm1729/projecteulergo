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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sqrd(x int) int {
	return x * x
}
