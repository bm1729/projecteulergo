package projecteulergo

import "testing"

func TestProblem3(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{13195, 29},
		{600851475143, 6857},
	}
	for _, c := range cases {
		got := Problem3(c.in)
		if got != c.want {
			t.Errorf("Problem3(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
