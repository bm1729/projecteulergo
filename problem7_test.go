package projecteulergo

import "testing"

func TestProblem7(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{6, 13},
		{10001, 104743},
	}
	for _, c := range cases {
		got := Problem7(c.in)
		if got != c.want {
			t.Errorf("Problem7(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
