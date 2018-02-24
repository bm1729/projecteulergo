package projecteulergo

// Problem6 - Sum square difference
func Problem6(x int) int {
	squareOfSum := x * x * (x + 1) * (x + 1) / 4
	sumOfSquare := x * (x + 1) * ((2 * x) + 1) / 6

	return squareOfSum - sumOfSquare
}
