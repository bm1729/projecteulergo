package projecteulergo

import (
	"strconv"
)

func isPalidromic(x int) bool {
	s := strconv.Itoa(x)
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}

	return true
}

// Problem4 - Largest palindrome product
func Problem4(x int) int {
	m := 0
	for i := 1; i < x; i++ {
		for j := 1; j < x; j++ {
			p := i * j
			if p > m && isPalidromic(p) {
				m = p
			}
		}
	}
	return m
}
