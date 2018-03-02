package projecteulergo

func pythagoreanTriple(a, b, c int) bool {
	return sqrd(a)+sqrd(b) == sqrd(c)
}

// Problem9 - Special Pythagorean triplet
func Problem9() int {
	for a := 1; a <= 1000; a++ {
		for b := a; b <= 1000; b++ {
			c := 1000 - a - b
			if pythagoreanTriple(a, b, c) {
				return a * b * c
			}
		}
	}
	panic("Can't find pythagorean triple")
}
