package projecteulergo

import "testing"

func TestProblem6(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 2640},
		{100, 25164150},
	}
	for _, c := range cases {
		got := Problem6(c.in)
		if got != c.want {
			t.Errorf("Problem6(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
