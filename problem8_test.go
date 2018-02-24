package projecteulergo

import "testing"

func TestProblem8(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{4, 5832},
		{13, 23514624000},
	}
	for _, c := range cases {
		got := Problem8(c.in)
		if got != c.want {
			t.Errorf("Problem8(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
