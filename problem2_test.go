package projecteulergo

import "testing"

func TestProblem2(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{100, 44},
		{4000000, 4613732},
	}
	for _, c := range cases {
		got := Problem2(c.in)
		if got != c.want {
			t.Errorf("Problem2(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
