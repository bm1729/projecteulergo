package projecteulergo

import "testing"

func TestProblem9(t *testing.T) {
	want := 1
	got := Problem9()
	if got != want {
		t.Errorf("Problem9() == %d, want %d", got, want)
	}
}
