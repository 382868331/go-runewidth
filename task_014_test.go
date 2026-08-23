package runewidth

import "testing"

func TestTask014WrapAllowsExactFit(t *testing.T) {
	for _, tc := range []struct {
		s string
		w int
	}{{"ab", 2}, {"界", 2}} {
		if got := NewCondition().Wrap(tc.s, tc.w); got != tc.s {
			t.Fatalf("got=%q want=%q", got, tc.s)
		}
	}
}
