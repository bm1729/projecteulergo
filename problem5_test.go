package projecteulergo

import "testing"

func TestProblem5(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 2520},
		{20, 232792560},
	}
	for _, c := range cases {
		got := Problem5(c.in)
		if got != c.want {
			t.Errorf("Problem5(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
