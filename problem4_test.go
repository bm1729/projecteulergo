package projecteulergo

import "testing"

func TestProblem4(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{100, 9009},
		{1000, 906609},
	}
	for _, c := range cases {
		got := Problem4(c.in)
		if got != c.want {
			t.Errorf("Problem4(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
