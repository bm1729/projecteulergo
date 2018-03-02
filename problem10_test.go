package projecteulergo

import "testing"

func TestProblem10(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 17},
		{2000000, 142913828922},
	}
	for _, c := range cases {
		got := Problem10(c.in)
		if got != c.want {
			t.Errorf("Problem10(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
