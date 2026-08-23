package runewidth

import "testing"

func TestTask006WideRuneWidth(t *testing.T) {
	c := &Condition{}
	for _, r := range []rune{'界', '語'} {
		if got := c.RuneWidth(r); got != 2 {
			t.Fatalf("rune=%U width=%d", r, got)
		}
	}
}
