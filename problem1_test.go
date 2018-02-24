package projecteulergo

import "testing"

func TestProblem1(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 23},
		{1000, 233168},
	}
	for _, c := range cases {
		got := Problem1(c.in)
		if got != c.want {
			t.Errorf("Problem1(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
